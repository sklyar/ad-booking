package application

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sklyar/ad-booking/backend/internal/application/config"
	"github.com/sklyar/ad-booking/backend/internal/infrastructure/db"
	"github.com/sklyar/ad-booking/backend/internal/server"
	"github.com/sklyar/go-transact"
	"github.com/sklyar/go-transact/adapters/transactstd"
	"log/slog"
	"net"
)

type App struct {
	logger     *slog.Logger
	httpServer *server.Server
}

func Bootstrap(ctx context.Context, logger *slog.Logger, cfgFile string) (*App, error) {
	cfg, err := config.New(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	_, dbHandler, err := initDB(ctx, cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("init db: %w", err)
	}

	repoContainer := newRepositoryContainer(dbHandler)
	serviceContainer := newServiceContainer(repoContainer)

	ln, err := net.Listen("tcp", cfg.Server.Addr())
	if err != nil {
		return nil, fmt.Errorf("listen: %w", err)
	}
	httpServer := server.New(logger, ln, serviceContainer.ContactPersonService)

	return &App{
		logger:     logger,
		httpServer: httpServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.httpServer.Run(ctx)
}

func initDB(ctx context.Context, dbConfig config.Database) (*db.TxManager, db.Handler, error) {
	sqlDB, err := sql.Open("pgx", dbConfig.DSN)
	if err != nil {
		return nil, nil, fmt.Errorf("open database: %w", err)
	}

	manager, handler, err := transact.NewManager(transactstd.Wrap(sqlDB))
	if err != nil {
		return nil, nil, fmt.Errorf("create transaction manager: %w", err)
	}

	if err := handler.Ping(ctx); err != nil {
		return nil, nil, fmt.Errorf("ping database: %w", err)
	}

	return manager, handler, nil
}
