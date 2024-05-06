package server

import (
	"net/http"

	"github.com/esgi-challenge/backend/internal/middleware"
	"github.com/gin-gonic/gin"

	userHttp "github.com/esgi-challenge/backend/internal/user/http"
	userRepo "github.com/esgi-challenge/backend/internal/user/repository"
	userUseCase "github.com/esgi-challenge/backend/internal/user/usecase"
)

func (s *Server) SetupHandlers() error {
	// Repo
	userRepo := userRepo.NewUserRepository(s.psqlDB)

	// UseCase
	userUseCase := userUseCase.NewUserUseCase(userRepo, s.cfg, s.logger)

	// Handlers
	userHandlers := userHttp.NewUserHandlers(userUseCase, s.cfg, s.logger)

	mw := middleware.InitMiddlewareManager(s.cfg, s.logger)

	s.router.Use(mw.RequestMiddleware())

	api := s.router.Group("/api")

	userGroup := api.Group("/users")
	userHttp.SetupUserRoutes(userGroup, userHandlers)

	health := api.Group("/healthz")
	health.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	return nil
}
