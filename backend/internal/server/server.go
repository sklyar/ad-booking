package server

import (
	"context"
	"github.com/sklyar/ad-booking/backend/api/gen/contactperson/contactpersonconnect"
	"github.com/sklyar/ad-booking/backend/internal/server/person"
	"github.com/sklyar/ad-booking/backend/internal/service"
	"log/slog"
	"net"
	"net/http"
)

type Server struct {
	ln     net.Listener
	mux    *http.ServeMux
	logger *slog.Logger
}

func New(
	logger *slog.Logger,
	ln net.Listener,
	contactPersonService service.Person,
) *Server {
	mux := http.NewServeMux()
	mount := func(path string, handler http.Handler) {
		mux.Handle(path, handler)
	}

	contactPersonHandler := person.New(contactPersonService)
	mount(contactpersonconnect.NewServiceHandler(contactPersonHandler))

	return &Server{
		ln:     ln,
		mux:    mux,
		logger: logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	srv := &http.Server{
		Handler: s.mux,
	}

	go func() {
		<-ctx.Done()

		if err := srv.Shutdown(context.Background()); err != nil {
			s.logger.Error("server shutdown", err)
		}
	}()

	s.logger.Info("starting server", slog.String("addr", s.ln.Addr().String()))

	return srv.Serve(s.ln)
}
