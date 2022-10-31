package server

import (
	"context"
	"errors"
	"net/http"

	"fx-test/api/config"
	"go.uber.org/zap"
)

type Route interface {
	http.Handler
	Pattern() string
}

type Server struct {
	*http.Server
	lgr *zap.Logger
}

func NewHTTPServer(cfg *config.APIConfig, mux *http.ServeMux, lgr *zap.Logger) *Server {
	return &Server{
		Server: &http.Server{
			Addr:    cfg.Host,
			Handler: mux,
		},
		lgr: lgr,
	}
}

func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, r := range routes {
		mux.Handle(r.Pattern(), r)
	}

	return mux
}

func (s *Server) StartServer() error {
	var err error
	if s.Addr == "" {
		err = errors.New("server host address cannot be nil")
		s.lgr.Error("Internal Server Error", zap.Error(err))
		return err
	}

	if s.Handler == nil {
		err = errors.New("server should have at least one handler")
		s.lgr.Error("Internal Server Error", zap.Error(err))
		return err
	}

	s.lgr.Info("Starting HTTP server at", zap.String("addr", s.Addr))
	go s.listenAndServe()
	s.lgr.Info("Successfully started HTTP server at", zap.String("addr", s.Addr))

	return err
}

func (s *Server) listenAndServe() {
	c := make(chan error)
	go func() {
		c <- s.ListenAndServe()
	}()

	err := <-c
	if err != nil {
		s.lgr.Error("Error occurred when listening at", zap.String("addr", s.Addr), zap.Error(err))
	}
}

func (s *Server) StopServer(ctx context.Context) error {
	s.lgr.Info("Shutting down HTTP server at", zap.String("addr", s.Addr))
	return s.Shutdown(ctx)
}
