package main

import (
	"github.com/Mortimor1/mikromon-core/internal/core"
	"github.com/Mortimor1/mikromon-core/pkg/logging"
	"github.com/gookit/config"
	"github.com/gookit/config/yaml"
)

func main() {
	logger := logging.GetLogger()

	loadConfig()
	server := new(core.Server)

	port, _ := config.String("port")
	if err := server.Run(port); err != nil {
		logger.Fatalf("error running http server: %s", err.Error())
	}
}

func loadConfig() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config/config.yml")
	if err != nil {
		panic(err)
	}
}
