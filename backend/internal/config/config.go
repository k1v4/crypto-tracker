package config

import (
	"backend/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Listen struct {
		Type   string `yaml:"type" env-default:"port"`
		Port   string `yaml:"port" env-default:"8080"`
		BindIp string `yaml:"bind_ip" env-default:"127.0.1"`
	} `yaml:"listen"`
	CMCApiKey string `yaml:"cmc_api_key"`
}

var once sync.Once
var instance *Config

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("reading application configuration")

		instance = &Config{}

		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			desc, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(desc)
			logger.Fatal(err)
		}
	})

	return instance
}
