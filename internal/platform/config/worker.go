package config

type workerConfig struct {
}

func ProvideWorkerConfig() Worker {
	return &workerConfig{}
}

func (w *workerConfig) GetMongoUrl() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w *workerConfig) GetMongoDatabase() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w *workerConfig) GetKafkaUrl() (string, error) {
	//TODO implement me
	panic("implement me")
}
