package server

import (
	"context"
	"log/slog"
	"net/http"
)

type Server struct {
	addr   string
	mux    *http.ServeMux
	logger *slog.Logger
}

func New(logger *slog.Logger, addr string) *Server {
	mux := http.NewServeMux()
	return &Server{
		addr:   addr,
		mux:    mux,
		logger: logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.mux,
	}

	go func() {
		<-ctx.Done()

		if err := srv.Shutdown(context.Background()); err != nil {
			s.logger.Error("server shutdown", err)
		}
	}()

	s.logger.Info("starting server", slog.String("addr", s.addr))

	return srv.ListenAndServe()
}
