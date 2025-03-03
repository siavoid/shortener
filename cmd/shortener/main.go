package main

import (
	"log"

	"github.com/siavoid/shortener/config"

	"github.com/joho/godotenv"

	"github.com/siavoid/shortener/internal/app/shortener"

	_ "github.com/siavoid/shortener/docs"
)

// @title URL Shortener API
// @version         2.0
// @description API for shortening URLs.

// @host localhost:8080
// @BasePath /

func main() {
	err := godotenv.Load()
	if err != nil {
		//log.Fatalf("Error loading .env file")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error read config")
		return
	}
	shortener.Run(cfg)
}
