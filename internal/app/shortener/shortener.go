package shortener

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/siavoid/shortener/config"
	v1 "github.com/siavoid/shortener/internal/controllers/http/v1"

	"github.com/siavoid/shortener/internal/repo/pgrepo"

	"github.com/siavoid/shortener/internal/repo/urlstore"
	"github.com/siavoid/shortener/internal/usecase"
	"github.com/siavoid/shortener/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	l.Debug("config: %+v\n\n", cfg)

	var repo usecase.URLStoreInterface
	if cfg.PG.URL != "" {
		l.Debug("pg: %s", cfg.PG.URL)
		db, err := pgrepo.NewPostgresRepo(cfg.PG.URL, l)
		if err != nil {
			l.Fatal("postgres connect : %s", err.Error())
			return
		}
		repo = db
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
		if err := server.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Создаем канал для получения сигналов
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // Подписываемся на сигнал прерывания (Ctrl+C)

	// Ожидаем сигнала
	<-signalChan
	log.Println("Received shutdown signal, stopping server...")

	// Останавливаем сервер
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Stop(ctx); err != nil {
		l.Error("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}
