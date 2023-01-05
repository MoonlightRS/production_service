package main

import (
	"log"
	"production_service/app/internal/config"
	"production_service/app/pkg/logging"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logger := logging.GetLogger()
	app, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}
}
