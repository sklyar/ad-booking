package apptest

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/sklyar/ad-booking/backend/internal/application/config"

	"github.com/google/uuid"
	"github.com/sklyar/ad-booking/backend/internal/infrastructure/database"
	"github.com/stretchr/testify/require"
)

var (
	sourceDB database.Handler
	once     sync.Once
)

// NewDB creates a new database handler and transaction manager for testing purposes.
// It clones the source database and returns the cloned database handler and transaction manager.
func NewDB(t *testing.T, ctx context.Context, logger *slog.Logger, cfg config.Database) *database.Database {
	t.Helper()

	baseDSN, err := url.Parse(cfg.DSN)
	if err != nil {
		t.Fatal(err, "failed to parse DSN")
	}

	sourceDBName := extractDBName(baseDSN)

	once.Do(func() {
		parsedDSN := *baseDSN
		parsedDSN.Path = "postgres"

		dbConfig := database.Config{
			DSN:      parsedDSN.String(),
			Logger:   logger,
			LogLevel: cfg.LogLevel,
		}
		db, err := database.New(ctx, dbConfig)
		require.NoError(t, err, "failed to create source database")

		sourceDB = db.Handler
	})

	clonedDBName := generateClonedDBName()
	clonedDSN := *baseDSN
	clonedDSN.Path = clonedDBName

	// Create and connect to the cloned database.
	clonedDB, err := cloneDB(ctx, logger, cfg.LogLevel, sourceDBName, clonedDBName, clonedDSN.String())
	if err != nil {
		t.Fatal(err, "failed to clone database")
	}

	t.Cleanup(func() {
		if err := clonedDB.Handler.Close(); err != nil {
			t.Logf("closing test database: %v", err)
		}

		query := fmt.Sprintf(`DROP DATABASE %q WITH (FORCE)`, clonedDBName)
		_, err := sourceDB.Exec(context.Background(), query)
		if err != nil {
			t.Logf("dropping test database: %v", err)
		}
	})

	return clonedDB
}

func cloneDB(ctx context.Context, logger *slog.Logger, logLevel, template, target, targetDSN string) (*database.Database, error) {
	query := fmt.Sprintf(`CREATE DATABASE %q WITH TEMPLATE %q`, target, template)
	if _, err := sourceDB.Exec(ctx, query); err != nil {
		return nil, fmt.Errorf("failed to clone database: %w", err)
	}

	dbConfig := database.Config{
		DSN:      targetDSN,
		Logger:   logger,
		LogLevel: logLevel,
	}
	db, err := database.New(ctx, dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the cloned database: %w", err)
	}
	return db, nil
}

func extractDBName(dsn *url.URL) string {
	dbName := dsn.Query().Get("dbname")
	if dbName == "" {
		parts := strings.Split(dsn.Path, "/")
		if len(parts) >= 2 {
			dbName = parts[1]
		}
	}
	return dbName
}

func generateClonedDBName() string {
	clonedDBName := "test_" + uuid.New().String()
	// PostgreSQL database names are limited to 63 characters.
	if len(clonedDBName) > 63 {
		clonedDBName = clonedDBName[:63]
	}
	return clonedDBName
}
