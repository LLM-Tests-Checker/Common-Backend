package main

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/workers/model_check"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Errorf("godotenv.Load: %s", err)
		os.Exit(1)
	}

	ctx := context.Background()
	logger := configureLogger(ctx)

	logger.Info("Worker is starting")

	worker := configureWorker(logger)

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
}

func configureWorker(logger *logrus.Logger) worker {
	modelCheckWorker := model_check.NewWorker(logger)

	return modelCheckWorker
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

type worker interface {
	Start(ctx context.Context) error
}
