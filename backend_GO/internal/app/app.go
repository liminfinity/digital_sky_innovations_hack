package app

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kurochkinivan/commit_history/internal/config"
	historyController "github.com/kurochkinivan/commit_history/internal/controller/http/v1/history"
	historyUseCase "github.com/kurochkinivan/commit_history/internal/usecase/history"
	"github.com/kurochkinivan/commit_history/internal/usecase/repository/github"
)

type App struct {
	log    *slog.Logger
	cfg    *config.Config
	server *http.Server
}

func New(log *slog.Logger, cfg *config.Config) *App {
	r := httprouter.New()

	storage := github.New(cfg.Github.AccessToken, cfg.Github.Owner, cfg.Github.Repo)
	usecase := historyUseCase.New(storage)

	log.Info("registering history controller")
	c := historyController.New(usecase)
	c.Register(r)

	log.Info("creating HTTP server",
		slog.String("host", cfg.HTTP.Host),
		slog.String("port", cfg.HTTP.Port),
	)

	server := &http.Server{
		Addr:         net.JoinHostPort(cfg.HTTP.Host, cfg.HTTP.Port),
		Handler:      r,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		IdleTimeout:  cfg.HTTP.IdleTimeout,
	}

	return &App{
		log:    log,
		cfg:    cfg,
		server: server,
	}
}

func (a *App) Start() error {
	const op = "app.Start"

	a.log.Info("creating listener",
		slog.String("host", a.cfg.HTTP.Host),
		slog.String("port", a.cfg.HTTP.Port),
	)
	l, err := net.Listen("tcp", net.JoinHostPort(a.cfg.HTTP.Host, a.cfg.HTTP.Port))
	if err != nil {
		a.log.Error("failed to create listener", slog.String("op", op), slog.Any("err", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info("starting server")
	err = a.server.Serve(l)
	if err != nil && err != http.ErrServerClosed {
		a.log.Error("server failed", slog.String("op", op), slog.Any("err", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info("server stopped gracefully")
	return nil
}
