package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
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
