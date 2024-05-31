package main

import (
	"github.com/BASSWAVE/RESTapi/internal/config"
	"github.com/BASSWAVE/RESTapi/internal/storage/sqlite"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting app")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", err)
		os.Exit(1)
	}

	id, err := storage.SaveURL("https://google.com", "google4")
	if err != nil {
		log.Error("failed to save", err)
	}
	log.Info("saved url", slog.Int64("id", id))

	url, err := storage.GetURL("google2")
	if err != nil {
		log.Error("failed to get", err)
	}
	log.Info("get url", slog.String("url", url))

	// TODO init router
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
