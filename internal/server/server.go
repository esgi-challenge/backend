package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	cfg    *config.Config
	logger logger.Logger
}

func NewServer(cfg *config.Config, logger logger.Logger) *Server {
	return &Server{
		engine: gin.New(),
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) Run() error {
	s.engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world")
	})

	server := &http.Server{
		Addr:    ":" + s.cfg.Port,
		Handler: s.engine,
	}

	go func() {
		s.logger.Info("Server listening on PORT: " + s.cfg.Port)
		if err := server.ListenAndServe(); err != nil {
			s.logger.Fatal("Error starting the server: ", err)
		}
	}()

	stopSignal := make(chan os.Signal)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)

	<-stopSignal
	s.logger.Info("Shutting down the server...")

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	s.logger.Info("Server shutdown complete.")

	return server.Shutdown(ctx)
}
