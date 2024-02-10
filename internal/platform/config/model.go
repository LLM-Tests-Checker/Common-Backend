package config

import "time"

type Environment int

const (
	EnvironmentLocal Environment = iota + 1
	EnvironmentProduction
)

type Server interface {
	GetServerPort() (string, error)
	GetEnvironment() (Environment, error)

	GetAccessTokenSecret() (string, error)
	GetRefreshTokenSecret() (string, error)
	GetTokenIssuer() (string, error)
	GetAccessTokenLifetime() (time.Duration, error)
	GetRefreshTokenLifetime() (time.Duration, error)

	GetMongoUrl() (string, error)
	GetMongoDatabase() (string, error)
}

type WorkerLaunchLLMCheck interface {
	GetEnvironment() (Environment, error)

	GetMongoUrl() (string, error)
	GetMongoDatabase() (string, error)

	GetKafkaUrl() (string, error)
	GetKafkaTopicLLMCheck() (string, error)
}

type WorkerDropOldInProgressLLMCheck interface {
	GetEnvironment() (Environment, error)
}

type ConsumerLLMResult interface {
	GetEnvironment() (Environment, error)

	GetMongoUrl() (string, error)
	GetMongoDatabase() (string, error)

	GetKafkaUrl() (string, error)
	GetKafkaTopicLLMResult() (string, error)
	GetConsumerGroup() (string, error)
}
