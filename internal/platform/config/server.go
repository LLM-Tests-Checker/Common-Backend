package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type serverConfig struct {
	portCached                 *string
	environmentCached          Environment
	accessTokenSecretCached    *string
	refreshTokenSecretCached   *string
	accessTokenLifetimeCached  time.Duration
	refreshTokenLifetimeCached time.Duration
	tokenIssuerCached          *string
	mongoUrlCached             *string
	mongoDbCached              *string
}

func ProvideServerConfig() Server {
	return &serverConfig{}
}

func (config *serverConfig) GetServerPort() (string, error) {
	if config.portCached != nil {
		return *config.portCached, nil
	}
	serverPort, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		return "", errors.New("environment variable \"SERVER_PORT\" not provided")
	}

	config.portCached = &serverPort
	return *config.portCached, nil
}

func (config *serverConfig) GetEnvironment() (Environment, error) {
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

func (config *serverConfig) GetAccessTokenSecret() (string, error) {
	if config.accessTokenSecretCached != nil {
		return *config.accessTokenSecretCached, nil
	}

	secret, exists := os.LookupEnv("TOKEN_ACCESS_SIGN_SECRET")
	if !exists {
		return "", errors.New("environment variable \"TOKEN_ACCESS_SIGN_SECRET\" not provided")
	}

	config.accessTokenSecretCached = &secret

	return *config.accessTokenSecretCached, nil
}

func (config *serverConfig) GetRefreshTokenSecret() (string, error) {
	if config.refreshTokenSecretCached != nil {
		return *config.refreshTokenSecretCached, nil
	}

	secret, exists := os.LookupEnv("TOKEN_REFRESH_SIGN_SECRET")
	if !exists {
		return "", errors.New("environment variable \"TOKEN_REFRESH_SIGN_SECRET\" not provided")
	}

	config.refreshTokenSecretCached = &secret

	return *config.refreshTokenSecretCached, nil
}

func (config *serverConfig) GetTokenIssuer() (string, error) {
	if config.tokenIssuerCached != nil {
		return *config.tokenIssuerCached, nil
	}

	issuer, exists := os.LookupEnv("TOKEN_ISSUER")
	if !exists {
		return "", errors.New("environment variable \"TOKEN_ISSUER\" not provided")
	}

	config.tokenIssuerCached = &issuer

	return *config.tokenIssuerCached, nil
}

func (config *serverConfig) GetAccessTokenLifetime() (time.Duration, error) {
	if config.accessTokenLifetimeCached != 0 {
		return config.accessTokenLifetimeCached, nil
	}

	lifetime, exists := os.LookupEnv("TOKEN_ACCESS_LIFETIME_SECONDS")
	if !exists {
		return 0, errors.New("environment variable \"TOKEN_ACCESS_LIFETIME_SECONDS\" not provided")
	}

	lifetimeSeconds, err := strconv.ParseInt(lifetime, 10, 0)
	if err != nil {
		return 0, err
	}

	lifetimeDuration := time.Duration(lifetimeSeconds) * time.Second
	config.accessTokenLifetimeCached = lifetimeDuration

	return config.accessTokenLifetimeCached, nil
}

func (config *serverConfig) GetRefreshTokenLifetime() (time.Duration, error) {
	if config.refreshTokenLifetimeCached != 0 {
		return config.refreshTokenLifetimeCached, nil
	}

	lifetime, exists := os.LookupEnv("TOKEN_REFRESH_LIFETIME_SECONDS")
	if !exists {
		return 0, errors.New("environment variable \"TOKEN_REFRESH_LIFETIME_SECONDS\" not provided")
	}

	lifetimeSeconds, err := strconv.ParseInt(lifetime, 10, 0)
	if err != nil {
		return 0, err
	}

	lifetimeDuration := time.Duration(lifetimeSeconds) * time.Second
	config.refreshTokenLifetimeCached = lifetimeDuration

	return config.refreshTokenLifetimeCached, nil
}

func (config *serverConfig) GetMongoUrl() (string, error) {
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

func (config *serverConfig) GetMongoDatabase() (string, error) {
	if config.mongoDbCached != nil {
		return *config.mongoDbCached, nil
	}

	database, exists := os.LookupEnv("MONGODB_DATABASE")
	if !exists {
		return "", errors.New("environment variable \"MONGODB_DATABASE\" not provided")
	}

	config.mongoDbCached = &database

	return *config.mongoDbCached, nil
}
