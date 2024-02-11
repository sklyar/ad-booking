package main

import (
	"context"
	"github.com/sklyar/ad-booking/backend/internal/application"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultConfigPath = "config.yaml"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfgPath := os.Getenv("ADB_CONFIG_PATH")
	if cfgPath == "" {
		cfgPath = defaultConfigPath
	}

	app, err := application.Bootstrap(ctx, logger, cfgPath)
	if err != nil {
		logger.Error("failed to bootstrap app", err)
		return
	}

	if err := app.Run(ctx); err != nil {
		logger.Error("failed to run app", err)
		return
	}
}
