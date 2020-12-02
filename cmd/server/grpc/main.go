package main

import (
	"fmt"
	cfg "github.com/product/internal/config"

	grpcAPI "github.com/product/pkg/infrastructure/grpc/api"
	"github.com/product/pkg/shared/logger"
	"log"
)

func main() {
	config := cfg.GetConfig()

	fmt.Println(config.Logger)
	logConfig := logger.Configuration{
		EnableConsole:     config.Logger.Console.Enable,
		ConsoleJSONFormat: config.Logger.Console.JSON,
		ConsoleLevel:      config.Logger.Console.Level,
		EnableFile:        config.Logger.File.Enable,
		FileJSONFormat:    config.Logger.File.JSON,
		FileLevel:         config.Logger.File.Level,
		FileLocation:      config.Logger.File.Path,
	}

	if err := logger.NewLogger(logConfig, logger.InstanceZapLogger); err != nil {
		log.Fatalf("Could not instantiate log %v", err)
	}

	grpcAPI.RunServer()
}
