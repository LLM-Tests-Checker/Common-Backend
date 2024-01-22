package main

import (
	"context"
	config2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/config"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/producers/llm_check"
	llm2 "github.com/LLM-Tests-Checker/Common-Backend/internal/storage/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/storage/test"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/workers/model_check"
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

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Errorf("godotenv.Load: %s", err)
		os.Exit(1)
	}

	ctx := context.Background()
	logger := configureLogger(ctx)
	config := config2.ProvideWorkerConfig()

	logger.Info("Worker is starting")

	worker, mongoClient, kafkaWriter := configureWorker(ctx, logger, config)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		logger.Info("Worker started")

		err := worker.Start(ctx)
		if err != nil {
			logger.Errorf("Worker returned error: %s", err)
			close(done)
		}
	}()

	<-done
	logger.Info("Worker is stopping")

	cancel()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)

	err = mongoClient.Disconnect(ctx)
	if err != nil {
		logger.Errorf("mongoClient.Disconnect: %s", err)
		os.Exit(1)
	}

	err = kafkaWriter.Close()
	if err != nil {
		logger.Errorf("kafkaWriter.Close: %s", err)
		os.Exit(1)
	}

	logger.Infof("Worker stopped")
}

func configureWorker(
	ctx context.Context,
	logger *logrus.Logger,
	config config2.Worker,
) (*model_check.Worker, *mongo.Client, *kafka.Writer) {
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
		SetAppName("worker").
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

	testsStorage := test.NewTestsStorage(logger, mongoDatabase)
	llmStorage := llm2.NewLLMStorage(logger, mongoDatabase)

	kafkaWriter := configureKafkaWriter(logger, config)

	llmCheckProducer := llm_check.NewProducer(logger, kafkaWriter)

	modelCheckWorker := model_check.NewWorker(logger, llmStorage, testsStorage, llmCheckProducer)

	return modelCheckWorker, mongoClient, kafkaWriter
}

func configureLogger(ctx context.Context) *logrus.Logger {
	logger := logrus.New()

	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.PrettyPrint = false

	launchEnvironment, exists := os.LookupEnv("ENVIRONMENT")
	if !exists {
		logger.Errorf("ENVIRONMENT enviroment not provided")
		os.Exit(1)
	}

	logger.WithContext(ctx)
	logger.SetReportCaller(true)
	logger.SetFormatter(formatter)
	logger.WithField("environment", launchEnvironment)
	logger.WithField("application", "worker")

	return logger
}

func configureKafkaWriter(
	logger *logrus.Logger,
	config config2.Worker,
) *kafka.Writer {
	topic, err := config.GetKafkaTopicLLMCheck()
	if err != nil {
		logger.Errorf("config.GetKafkaTopicLLMCheck: %s", err)
		os.Exit(1)
	}

	kafkaUrl, err := config.GetKafkaUrl()
	if err != nil {
		logger.Errorf("config.GetKafkaUrl: %s", err)
		os.Exit(1)
	}

	writer := kafka.Writer{
		Addr:                   kafka.TCP(kafkaUrl),
		Topic:                  topic,
		MaxAttempts:            3,
		BatchTimeout:           0,
		ReadTimeout:            5 * time.Second,
		WriteTimeout:           5 * time.Second,
		RequiredAcks:           kafka.RequireOne,
		Async:                  false,
		Logger:                 nil,
		ErrorLogger:            nil,
		Transport:              nil,
		AllowAutoTopicCreation: true,
	}

	return &writer
}
