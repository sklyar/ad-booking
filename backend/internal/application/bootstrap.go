package application

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sklyar/ad-booking/backend/internal/application/config"
	"github.com/sklyar/ad-booking/backend/internal/server"
	"github.com/sklyar/go-transact"
	"github.com/sklyar/go-transact/adapters/transactstd"
	"github.com/sklyar/go-transact/txsql"
	"log/slog"
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

	httpServer := server.New(logger, cfg.Server.Addr())

	txManager, db, err := initDB(ctx, cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("init db: %w", err)
	}

	_ = txManager
	_ = db

	return &App{
		logger:     logger,
		httpServer: httpServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.httpServer.Run(ctx)
}

func initDB(ctx context.Context, dbConfig config.Database) (*transact.Manager, txsql.DB, error) {
	sqlDB, err := sql.Open("pgx", dbConfig.DSN)
	if err != nil {
		return nil, nil, fmt.Errorf("open database: %w", err)
	}

	manager, db, err := transact.NewManager(transactstd.Wrap(sqlDB))
	if err != nil {
		return nil, nil, fmt.Errorf("create transaction manager: %w", err)
	}

	if err := db.Ping(ctx); err != nil {
		return nil, nil, fmt.Errorf("ping database: %w", err)
	}

	return manager, db, nil
}
