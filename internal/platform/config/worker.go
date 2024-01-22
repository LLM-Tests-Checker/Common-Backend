package config

import (
	"errors"
	"fmt"
	"os"
)

type workerConfig struct {
	environmentCached        Environment
	mongoUrlCached           *string
	mongoDatabaseCached      *string
	kafkaUrlCached           *string
	kafkaTopicLLMCheckCached *string
}

func ProvideWorkerConfig() Worker {
	return &workerConfig{}
}

func (config *workerConfig) GetEnvironment() (Environment, error) {
	if config.environmentCached != 0 {
		return config.environmentCached, nil
	}

	environment, exists := os.LookupEnv("ENVIRONMENT")
	if !exists {
		return 0, errors.New("environment variable \"ENVIRONMENT\" not provided")
	}

	switch environment {
	case "LOCAL":
		config.environmentCached = EnvironmentLocal
	case "PRODUCTION":
		config.environmentCached = EnvironmentProduction
	default:
		return 0, errors.New(fmt.Sprintf("unknown value of variable \"ENVIRONMENT\": %s", environment))
	}

	return config.environmentCached, nil
}

func (config *workerConfig) GetMongoUrl() (string, error) {
	if config.mongoUrlCached != nil {
		return *config.mongoUrlCached, nil
	}

	url, exists := os.LookupEnv("MONGODB_URL")
	if !exists {
		return "", errors.New("environment variable \"MONGODB_URL\" not provided")
	}

	config.mongoUrlCached = &url

	return *config.mongoUrlCached, nil
}

func (config *workerConfig) GetMongoDatabase() (string, error) {
	if config.mongoDatabaseCached != nil {
		return *config.mongoDatabaseCached, nil
	}

	database, exists := os.LookupEnv("MONGODB_DATABASE")
	if !exists {
		return "", errors.New("environment variable \"MONGODB_DATABASE\" not provided")
	}

	config.mongoDatabaseCached = &database

	return *config.mongoDatabaseCached, nil
}

func (config *workerConfig) GetKafkaUrl() (string, error) {
	if config.kafkaUrlCached != nil {
		return *config.kafkaUrlCached, nil
	}

	kafkaUrl, exists := os.LookupEnv("KAFKA_URL")
	if !exists {
		return "", errors.New("environment variable \"KAFKA_URL\" not provided")
	}

	config.kafkaUrlCached = &kafkaUrl

	return *config.kafkaUrlCached, nil
}

func (config *workerConfig) GetKafkaTopicLLMCheck() (string, error) {
	if config.kafkaTopicLLMCheckCached != nil {
		return *config.kafkaTopicLLMCheckCached, nil
	}

	kafkaTopic, exists := os.LookupEnv("KAFKA_TOPIC_LLM_CHECK")
	if !exists {
		return "", errors.New("environment variable \"KAFKA_TOPIC_LLM_CHECK\" not provided")
	}

	config.kafkaTopicLLMCheckCached = &kafkaTopic

	return *config.kafkaTopicLLMCheckCached, nil
}
