package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/siavoid/shortener/config"

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
		log.Println("Error loading .env file")
	}

	// Определение флагов
	address := flag.String("a", "", "Address to start the HTTP server (e.g., localhost:8888)")
	baseURL := flag.String("b", "", "Base URL for the shortened URL (e.g., http://localhost:8000)")

	// Вывод справочной информации
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Парсинг флагов
	flag.Parse()

	// Валидация адреса HTTP-сервера
	if *address != "" {
		if _, err := url.ParseRequestURI("http://" + *address); err != nil {
			log.Fatalf("Invalid address: %v", err)
		}
	}

	// Валидация базового URL
	if *baseURL != "" {
		// Проверка корректности базового URL
		if _, err := url.ParseRequestURI(*baseURL); err != nil {
			log.Fatalf("Invalid base URL: %v", err)
		}
	}

	cfg, err := config.NewConfig(*address, *baseURL)
	if err != nil {
		log.Fatalf("Error read config")
		return
	}

	shortener.Run(cfg)
}
