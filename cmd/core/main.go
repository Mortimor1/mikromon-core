package main

import (
	"github.com/Mortimor1/mikromon-core/internal/core"
	"github.com/gookit/config"
	"github.com/gookit/config/yaml"
	"log"
)

func main() {
	loadConfig()
	server := new(core.Server)

	port, _ := config.String("port")
	if err := server.Run(port); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
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
