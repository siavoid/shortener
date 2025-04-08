package shortener

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/siavoid/shortener/config"
	v1 "github.com/siavoid/shortener/internal/controllers/http/v1"
	"github.com/siavoid/shortener/internal/repo/pgrepo"
	"github.com/siavoid/shortener/internal/repo/urlstore"
	"github.com/siavoid/shortener/internal/usecase"
	"github.com/siavoid/shortener/pkg/logger"
	"github.com/siavoid/shortener/pkg/postgres"
)

func Run(cfg *config.Config) {

	l := logger.New(cfg.Log.Level)
	l.Debug("config: %+v\n\n", cfg)

	var repo usecase.URLStoreInterface
	if cfg.PG.URL != "" {
		db, err := postgres.New(cfg.PG.URL)
		if err != nil {
			l.Fatal("postgres connect : %s", err.Error())
			return
		}
		repo = pgrepo.New(db, l)
	} else {
		repostre, err := urlstore.NewURLStore(cfg.Repo.FileStore)
		if err != nil {
			l.Fatal("urlStore err : %w", err)
			return
		}
		repo = repostre
	}

	u := usecase.New(cfg, l, repo)
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
