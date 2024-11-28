package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/pkg/log"
)

type Server struct {
	*gin.Engine
	httpSrv *http.Server
	host    string
	port    int
	logger  *log.Logger
}

type Option func(*Server)

func NewServer(engine *gin.Engine, logger *log.Logger, opts ...Option) *Server {
	srv := &Server{
		Engine: engine,
		logger: logger,
	}
	for _, opt := range opts {
		opt(srv)
	}
	return srv
}

func WithHost(host string) Option {
	return func(srv *Server) {
		srv.host = host
	}
}

func WithPort(port int) Option {
	return func(srv *Server) {
		srv.port = port
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: s.Engine,
	}
	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.ErrorContext(ctx, fmt.Sprintf("listen: %s\n", err))
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.InfoContext(ctx, "Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpSrv.Shutdown(ctx); err != nil {
		s.logger.ErrorContext(ctx, "Server forced to shutdown: ", "err:", err.Error())
	}
	s.logger.InfoContext(ctx, "Server exiting")
	return nil
}
