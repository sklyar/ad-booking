package apptest

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/sklyar/ad-booking/backend/internal/application"
	"github.com/sklyar/ad-booking/backend/internal/application/config"
	"github.com/stretchr/testify/require"
)

const configPath = "config.yaml"
const suiteInitTimeout = 5 * time.Second
const logLevel = slog.LevelError

// Suite represents a test suite that includes the application instance and the HTTP client.
type Suite struct {
	app *application.App

	serverBaseURL string
	client        *http.Client
}

// NewSuite creates a new test suite for testing purposes.
func NewSuite(t *testing.T, ctx context.Context) *Suite {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	ctx, cancel := context.WithTimeout(ctx, suiteInitTimeout)

	cfgPath := resolveFilePath(t, configPath)

	cfg, err := config.New(cfgPath)
	if err != nil {
		t.Fatal(err, "failed to load config")
	}

	serverListener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	app, err := application.Bootstrap(
		ctx,
		cfg,
		application.WithLogger(logger),
		application.WithDB(NewDB(t, ctx, logger, cfg.Database)),
		application.WithServerListener(serverListener),
	)
	require.NoError(t, err)
	go func() {
		if err := app.Run(ctx); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			t.Error(err, "failed to run app")
		}
	}()

	s := &Suite{
		app:           app,
		serverBaseURL: serverURL(t, serverListener),
		client:        http.DefaultClient,
	}

	if !s.waitForServerReadiness(t, ctx, s.serverBaseURL) {
		cancel()
		t.Fatal("timeout exceeded waiting for server readiness")
	}

	t.Cleanup(cancel)

	return s
}

// resolveFilePath returns the absolute path to the file located in the same directory as the test file.
func resolveFilePath(t *testing.T, path string) string {
	t.Helper()

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot resolve current file path")
	}

	currentDir := filepath.Dir(currentFile)
	absolutePath := filepath.Join(currentDir, path)

	return absolutePath
}

// serverURL returns the URL of the given listener that use ipv4 loopback address.
func serverURL(t *testing.T, ln net.Listener) string {
	t.Helper()

	addr := ln.Addr().String()
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		t.Fatalf("failed to parse listener address: %v", err)
	}
	return fmt.Sprintf("http://127.0.0.1:%s", port)
}

// waitForServerReadiness waits for the server to be ready to accept connections.
func (s *Suite) waitForServerReadiness(t *testing.T, ctx context.Context, serverURL string) bool {
	t.Helper()

	healthURL, err := url.JoinPath(serverURL, "/health")
	require.NoError(t, err)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, healthURL, nil)
	require.NoError(t, err)

	const (
		maxAttempts = 10
		retryDelay  = 10 * time.Millisecond
	)

	for i := 0; i < maxAttempts; i++ {
		select {
		case <-ctx.Done():
			return false
		default:
			resp, err := s.client.Do(req)
			if err == nil && resp.StatusCode == http.StatusOK {
				return true
			}
			time.Sleep(retryDelay)
		}
	}

	return false
}
