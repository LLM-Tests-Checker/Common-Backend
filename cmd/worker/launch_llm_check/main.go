package main

import (
	"context"
	config2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/config"
	logger2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/producers/llm_check"
	llm2 "github.com/LLM-Tests-Checker/Common-Backend/internal/storage/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/storage/test"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/workers/launch_llm_check"
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

const applicationName = "worker_launch_llm_check"

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Errorf("godotenv.Load: %s", err)
		os.Exit(1)
	}

	config := config2.ProvideWorkerLaunchLLMCheckConfig()
	ctx := context.Background()
	logger := configureLogger(ctx, config)

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
	logger logger2.Logger,
	config config2.WorkerLaunchLLMCheck,
) (*launch_llm_check.Worker, *mongo.Client, *kafka.Writer) {
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

	testsStorage := test.NewTestsStorage(logger, mongoDatabase)
	llmStorage := llm2.NewLLMStorage(logger, mongoDatabase)

	kafkaWriter := configureKafkaWriter(logger, config)

	llmCheckProducer := llm_check.NewProducer(logger, kafkaWriter)

	modelCheckWorker := launch_llm_check.NewWorker(logger, llmStorage, testsStorage, llmCheckProducer)

	return modelCheckWorker, mongoClient, kafkaWriter
}

func configureLogger(
	ctx context.Context,
	config config2.WorkerLaunchLLMCheck,
) logger2.Logger {
	logger := logrus.New()

	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.PrettyPrint = false

	launchEnvironment, err := config.GetEnvironment()
	if err != nil {
		logger.Errorf("config.GetEnvironemnt: %s", err)
		os.Exit(1)
	}

	logger.SetReportCaller(true)
	logger.SetFormatter(formatter)

	return logger.WithFields(
		logrus.Fields{
			"environment": launchEnvironment,
			"application": applicationName,
		},
	).WithContext(ctx)
}

func configureKafkaWriter(
	logger logger2.Logger,
	config config2.WorkerLaunchLLMCheck,
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
		MaxAttempts:            2,
		ReadTimeout:            2 * time.Second,
		WriteTimeout:           2 * time.Second,
		RequiredAcks:           kafka.RequireOne,
		Async:                  false,
		Logger:                 kafka.LoggerFunc(logger.Infof),
		ErrorLogger:            kafka.LoggerFunc(logger.Errorf),
		AllowAutoTopicCreation: true,
	}

	return &writer
}
