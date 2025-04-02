package shortener

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/siavoid/shortener/config"
	v1 "github.com/siavoid/shortener/internal/controllers/http/v1"
	"github.com/siavoid/shortener/internal/repo/urlstore"
	"github.com/siavoid/shortener/internal/usecase"
	"github.com/siavoid/shortener/pkg/logger"
)

func Run(cfg *config.Config) {

	l := logger.New(cfg.Log.Level)

	var tempDB interface{}
	urlStore, err := urlstore.NewURLStore(cfg.Repo.FileStore)
	if err != nil {
		l.Fatal("urlStore err : %w", err)
	}
	u := usecase.New(cfg, l, tempDB, urlStore)
	server := v1.New(cfg, u, l)

	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Создаем канал для получения сигналов
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt) // Подписываемся на сигнал прерывания (Ctrl+C)

	// Ожидаем сигнала
	<-signalChan
	log.Println("Received shutdown signal, stopping server...")

	// Останавливаем сервер
	if err := server.Stop(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}
