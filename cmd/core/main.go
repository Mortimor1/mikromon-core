package main

import (
	"github.com/Mortimor1/mikromon-core/internal/config"
	"github.com/Mortimor1/mikromon-core/internal/core"
	"github.com/Mortimor1/mikromon-core/pkg/logging"
)

func main() {
	logger := logging.GetLogger()

	cfg := config.GetConfig()
	server := new(core.Server)

	if err := server.Run(cfg); err != nil {
		logger.Fatalf("error running http server: %s", err.Error())
	}
}
