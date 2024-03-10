package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/tracelog"
	pgxlogger "github.com/mcosta74/pgx-slog"
	"github.com/sklyar/go-transact"
	"github.com/sklyar/go-transact/adapters/transactstd"
	"github.com/sklyar/go-transact/txsql"
)

// TxManager is an alias for transact.Manager.
type TxManager = transact.Manager

// Handler is an alias for txsql.DB.
type Handler = txsql.DB

// Row is an alias for txsql.Row.
type Row = txsql.Row

// Config is a configuration for the database.
type Config struct {
	DSN      string
	Logger   *slog.Logger
	LogLevel string
}

// Database is a wrapper for the database handler and transaction manager.
type Database struct {
	Handler   Handler
	TxManager *TxManager
}

// New creates a new database handler and transaction manager.
func New(ctx context.Context, cfg Config) (*Database, error) {
	connConfig, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("parse Database config: %w", err)
	}

	tracer := &tracelog.TraceLog{Logger: pgxlogger.NewLogger(cfg.Logger), LogLevel: tracelog.LogLevelError}
	if cfg.LogLevel != "" {
		tracer.LogLevel, err = tracelog.LogLevelFromString(cfg.LogLevel)
		if err != nil {
			return nil, fmt.Errorf("parse log level: %w", err)
		}
	}
	connConfig.ConnConfig.Tracer = tracer

	pool, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	sqlDB := stdlib.OpenDBFromPool(pool)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	manager, handler, err := transact.NewManager(transactstd.Wrap(sqlDB))
	if err != nil {
		return nil, fmt.Errorf("create transaction manager: %w", err)
	}

	return &Database{
		Handler:   handler,
		TxManager: manager,
	}, nil
}
