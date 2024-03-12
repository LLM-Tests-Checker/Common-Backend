package main

import (
	"context"
	config2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/config"
	logger2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	llm2 "github.com/LLM-Tests-Checker/Common-Backend/internal/storage/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/workers/drop_old_in_progress_llm_check"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const applicationName = "worker_drop_old_in_progress_llm_check"

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnf("godotenv.Load: %s", err)
	}

	config := config2.ProvideWorkerDropOldLLMCheckConfig()
	ctx := context.Background()
	logger := configureLogger(ctx, config)

	logger.Info("Worker is starting")

	worker, mongoClient := configureWorker(ctx, logger, config)

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

	logger.Infof("Worker stopped")
}

func configureWorker(
	ctx context.Context,
	logger logger2.Logger,
	config config2.WorkerDropOldInProgressLLMCheck,
) (*drop_old_in_progress_llm_check.Worker, *mongo.Client) {
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

	worker := drop_old_in_progress_llm_check.NewWorker(logger, llmStorage)

	return worker, mongoClient
}

func configureLogger(
	ctx context.Context,
	config config2.WorkerDropOldInProgressLLMCheck,
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
