package server

import (
	"net/http"

	"github.com/esgi-challenge/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) SetupHandlers() error {
	mw := middleware.InitMiddlewareManager(s.cfg, s.logger)

	s.router.Use(mw.RequestMiddleware())

	api := s.router.Group("/api")

	health := api.Group("/healthz")

	health.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	return nil
}
