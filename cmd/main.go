package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/core/config"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

var logger *zap.Logger

func initLogger () {
	var err error

	logger, err = zap.NewDevelopment()

	if err != nil {
		log.Fatal("failed to start logger: ", err.Error())
	}
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return err
	}

	if err := envconfig.Process("app", &config.App); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func cleanupLogger(undoFn func()) {
	_ = logger.Sync()

	undoFn()
}	

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// initig logger, replacing global logger and defering its cleanup
	initLogger()
	undo := zap.ReplaceGlobals(logger)
	defer cleanupLogger(undo)


	
}
