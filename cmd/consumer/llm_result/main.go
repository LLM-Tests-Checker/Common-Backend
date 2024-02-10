package main

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/consumers/llm_result"
	config2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/config"
	logger2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	llm2 "github.com/LLM-Tests-Checker/Common-Backend/internal/storage/llm"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const applicationName = "consumer_llm_result"

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Errorf("godotenv.Load: %s", err)
		os.Exit(1)
	}

	config := config2.ProvideConsumerConfig()
	ctx := context.Background()
	logger := configureLogger(ctx)

	logger.Info("ConsumerLLMResult is starting")

	consumer, mongoClient, kafkaReader := configureConsumer(ctx, logger, config)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		logger.Info("ConsumerLLMResult started")

		err := consumer.Start(ctx)
		if err != nil {
			logger.Errorf("ConsumerLLMResult returned error: %s", err)
			close(done)
		}
	}()

	<-done
	logger.Info("ConsumerLLMResult is stopping")

	cancel()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)

	err = mongoClient.Disconnect(ctx)
	if err != nil {
		logger.Errorf("mongoClient.Disconnect: %s", err)
		os.Exit(1)
	}

	err = kafkaReader.Close()
	if err != nil {
		logger.Errorf("kafkaWriter.Close: %s", err)
		os.Exit(1)
	}

	logger.Infof("WorkerLaunchLLMCheck stopped")
}

func configureConsumer(
	ctx context.Context,
	logger logger2.Logger,
	config config2.ConsumerLLMResult,
) (*llm_result.Consumer, *mongo.Client, *kafka.Reader) {
	launchEnvironment, err := config.GetEnvironment()
	if err != nil {
		logger.Errorf("config.GetEnvironment: %s", err)
		os.Exit(1)
	}

	mongoUrl, err := config.GetMongoUrl()
	if err != nil {
		logger.Errorf("config.GetMongoUrl: %s", err)
		os.Exit(1)
	}

	mongodbLogLevel := options2.LogLevelInfo
	if launchEnvironment == config2.EnvironmentLocal {
		mongodbLogLevel = options2.LogLevelDebug
	}

	mongoLogOptions := options2.Logger().SetComponentLevel(options2.LogComponentAll, mongodbLogLevel)
	options := options2.Client().
		ApplyURI(mongoUrl).
		SetTimeout(time.Second).
		SetAppName(applicationName).
		SetConnectTimeout(10 * time.Second).
		SetMaxConnecting(10).
		SetMinPoolSize(5).
		SetRetryReads(true).
		SetMaxConnIdleTime(30 * time.Second).
		SetServerSelectionTimeout(10 * time.Second).
		SetLoggerOptions(mongoLogOptions)

	mongoClient, err := mongo.Connect(ctx, options)
	if err != nil {
		logger.Errorf("Can't connect to mongo: %s", err)
		os.Exit(1)
	}

	databaseName, err := config.GetMongoDatabase()
	if err != nil {
		logger.Errorf("config.GetMongoDatabase: %s", err)
		os.Exit(1)
	}

	mongoDatabase := mongoClient.Database(databaseName)

	llmStorage := llm2.NewLLMStorage(logger, mongoDatabase)
	kafkaReader := configureKafkaReader(logger, config)

	consumer := llm_result.NewConsumer(logger, kafkaReader, llmStorage)

	return consumer, mongoClient, kafkaReader
}

func configureLogger(ctx context.Context) logger2.Logger {
	logger := logrus.New()

	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.PrettyPrint = false

	launchEnvironment, exists := os.LookupEnv("ENVIRONMENT")
	if !exists {
		logger.Errorf("ENVIRONMENT enviroment not provided")
		os.Exit(1)
	}

	logger = logger.WithContext(ctx).Logger
	logger.SetReportCaller(true)
	logger.SetFormatter(formatter)

	return logger.WithFields(
		logrus.Fields{
			"environment": launchEnvironment,
			"application": applicationName,
		},
	)
}

func configureKafkaReader(
	logger logger2.Logger,
	config config2.ConsumerLLMResult,
) *kafka.Reader {
	topic, err := config.GetKafkaTopicLLMResult()
	if err != nil {
		logger.Errorf("config.GetKafkaTopicLLMResult: %s", err)
		os.Exit(1)
	}

	kafkaUrl, err := config.GetKafkaUrl()
	if err != nil {
		logger.Errorf("config.GetKafkaUrl: %s", err)
		os.Exit(1)
	}

	consumerGroup, err := config.GetConsumerGroup()
	if err != nil {
		logger.Errorf("config.GetConsumerGroup: %s", err)
		os.Exit(1)
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           []string{kafkaUrl},
		GroupID:           consumerGroup,
		Topic:             topic,
		MaxWait:           5 * time.Second,
		ReadBatchTimeout:  2 * time.Second,
		HeartbeatInterval: 1 * time.Second,
		SessionTimeout:    10 * time.Second,
		StartOffset:       kafka.FirstOffset,
		Logger:            kafka.LoggerFunc(logger.Infof),
		ErrorLogger:       kafka.LoggerFunc(logger.Errorf),
		IsolationLevel:    kafka.ReadCommitted,
		MaxAttempts:       2,
	})

	return reader
}
