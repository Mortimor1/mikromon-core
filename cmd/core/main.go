package main

import (
	"fmt"
	core "github.com/Mortimor1/mikromon-core"
	"github.com/gookit/config"
	"github.com/gookit/config/yaml"
	"log"
)

func main() {
	loadConfig()
	server := new(core.Server)
	if err := server.Run("8080"); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}

func loadConfig() {
	config.WithOptions(config.ParseEnv)

	// add driver for support yaml content
	config.AddDriver(yaml.Driver)
	// config.SetDecoder(config.Yaml, yaml.Decoder)

	err := config.LoadFiles("config/config.yml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("config data: \n %#v\n", config.Data())

	// load more files
	err = config.LoadFiles("config/config.yml")
	// can also load multi at once
	// err := config.LoadFiles("testdata/yml_base.yml", "testdata/yml_other.yml")
	if err != nil {
		panic(err)
	}
}
