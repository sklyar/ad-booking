package application

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/sklyar/ad-booking/backend/internal/application/config"
	"github.com/sklyar/ad-booking/backend/internal/infrastructure/database"
	"github.com/sklyar/ad-booking/backend/internal/server"
)

type AppOption func(app *App)

func WithLogger(logger *slog.Logger) AppOption {
	return func(app *App) {
		app.logger = logger
	}
}

func WithDB(db *database.Database) AppOption {
	return func(app *App) {
		app.db = db
	}
}

func WithServerListener(ln net.Listener) AppOption {
	return func(app *App) {
		app.httpServerListener = ln
	}
}

type App struct {
	ServiceContainer *ServiceContainer

	db                 *database.Database
	logger             *slog.Logger
	httpServer         *server.Server
	httpServerListener net.Listener
}

func Bootstrap(ctx context.Context, cfg *config.Config, opts ...AppOption) (*App, error) {
	var app App
	for _, opt := range opts {
		opt(&app)
	}

	logger := app.logger
	if logger == nil {
		var level slog.Level
		if err := level.UnmarshalText([]byte(cfg.Logger.Level)); err != nil {
			return nil, fmt.Errorf("parse log level: %w", err)
		}
		opts := &slog.HandlerOptions{
			Level: level,
		}
		logger = slog.New(slog.NewTextHandler(os.Stdout, opts))
	}

	if app.db == nil {
		var err error
		dbConfig := database.Config{
			DSN:      cfg.Database.DSN,
			Logger:   logger,
			LogLevel: cfg.Database.LogLevel,
		}
		app.db, err = database.New(ctx, dbConfig)
		if err != nil {
			return nil, fmt.Errorf("init Database: %w", err)
		}
	}

	repoContainer := newRepositoryContainer(app.db.Handler)
	serviceContainer := newServiceContainer(repoContainer)
	app.ServiceContainer = serviceContainer

	if app.httpServerListener == nil {
		ln, err := net.Listen("tcp", cfg.Server.Addr())
		if err != nil {
			return nil, fmt.Errorf("listen: %w", err)
		}
		app.httpServerListener = ln
	}
	app.httpServer = server.New(logger, app.httpServerListener, serviceContainer.PersonService)

	return &app, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.httpServer.Run(ctx)
}
