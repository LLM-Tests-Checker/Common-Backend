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

type Worker interface {
	GetEnvironment() (Environment, error)

	GetMongoUrl() (string, error)
	GetMongoDatabase() (string, error)

	GetKafkaUrl() (string, error)
}

type Consumer interface {
	GetMongoUrl() (string, error)
	GetMongoDatabase() (string, error)

	GetKafkaUrl() (string, error)
}
