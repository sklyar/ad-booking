package test

import (
	"context"
	"github.com/sklyar/ad-booking/backend/internal/application"
	"github.com/stretchr/testify/require"
	"log/slog"
	"os"
	"testing"
)

type Suite struct {
	app *application.App
}

func NewSuite(t *testing.T, ctx context.Context) *Suite {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	ctx, cancel := context.WithCancel(ctx)

	app, err := application.Bootstrap(ctx, logger, "./config.yaml")
	require.NoError(t, err)

	t.Cleanup(cancel)

	return &Suite{
		app: app,
	}
}
