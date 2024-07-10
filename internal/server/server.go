package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/pkg/gmap"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	ctxTimeout = 0
)

type Server struct {
	router         *gin.Engine
	cfg            *config.Config
	psqlDB         *gorm.DB
	logger         logger.Logger
	gmapApiManager *gmap.GmapApiManager
	storage        *storage.Storage
}

func NewServer(cfg *config.Config, psqlDB *gorm.DB, logger logger.Logger, gmapApiManager *gmap.GmapApiManager, storage *storage.Storage) *Server {
	return &Server{
		router:         gin.New(),
		cfg:            cfg,
		psqlDB:         psqlDB,
		logger:         logger,
		gmapApiManager: gmapApiManager,
		storage:        storage,
	}
}

func (s *Server) Run() error {
	s.logger.Info("Server: Setting up handlers...")
	if err := s.SetupHandlers(); err != nil {
		return err
	}
	s.logger.Info("Server: Handlers set")

	server := &http.Server{
		Addr:    s.cfg.BaseUrl + ":" + s.cfg.Port,
		Handler: s.router,
	}

	go func() {
		s.logger.Info("Server: Listening on " + s.cfg.BaseUrl + ":" + s.cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("Server: %s", err)
		}
	}()

	// Create channel to listen for termination events
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Channel in read only
	<-quit
	s.logger.Info("Server: Shutting down...")

	// Creating context with <ctxTimeout> seconds timeout, after that all app operations will be canceled
	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	// Wait for context timeout before shutting down
	select {
	case <-ctx.Done():
		s.logger.Info("Server: Shutdown complete")
	}

	return server.Shutdown(ctx)
}
