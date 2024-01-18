package config

type consumerConfig struct {
}

func ProvideConsumerConfig() Consumer {
	return &consumerConfig{}
}

func (c *consumerConfig) GetMongoUrl() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *consumerConfig) GetMongoDatabase() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *consumerConfig) GetKafkaUrl() (string, error) {
	//TODO implement me
	panic("implement me")
}
