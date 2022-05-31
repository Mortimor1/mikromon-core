package config

import (
	"github.com/Mortimor1/mikromon-core/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Debug bool `yaml:"debug"`
	Http  struct {
		BindIp string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"http"`
	Database struct {
		Mongo struct {
			Url string `yaml:"url"`
		} `yaml:"mongo"`
	} `yaml:"database"`
}

var instance *Config
var once sync.Once

var configPaths = []string{"config/core.yml", "core.yml", "/etc/mikromon/core.yml"}

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Read Config")
		instance = &Config{}

		var errorRead error
		for _, path := range configPaths {
			errorRead = cleanenv.ReadConfig(path, instance)
			if errorRead == nil {
				logger.Infof("Config loaded success path: %s", path)
				break
			}
		}

		if errorRead != nil {
			logger.Fatal(errorRead)
		}
	})
	return instance
}
