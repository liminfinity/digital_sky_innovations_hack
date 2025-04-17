package main

import (
	"log/slog"
	"os"

	"github.com/kurochkinivan/commit_history/internal/app"
	"github.com/kurochkinivan/commit_history/internal/config"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("config loaded", slog.String("env", cfg.Env))
	log.Info("creating application")

	application := app.New(log, cfg)

	log.Info("starting application")
	err := application.Start()
	if err != nil {
		log.Error("application stopped with error", slog.Any("err", err))
		panic(err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)
	}

	log.Info("logger initialized", slog.String("env", env))
	return log
}
