package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sklyar/ad-booking/backend/internal/application"
	"github.com/sklyar/ad-booking/backend/internal/application/config"
)

const (
	defaultConfigPath = "config.yaml"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfgPath := os.Getenv("ADB_CONFIG_PATH")
	if cfgPath == "" {
		cfgPath = defaultConfigPath
	}

	cfg, err := config.New(cfgPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app, err := application.Bootstrap(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to bootstrap app: %v", err)
		return
	}

	if err := app.Run(ctx); err != nil {
		log.Fatalf("failed to run app: %v", err)
		return
	}
}
