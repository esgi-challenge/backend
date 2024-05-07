package server

import (
	"net/http"

	_ "github.com/esgi-challenge/backend/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/esgi-challenge/backend/internal/middleware"
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
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := s.router.Group("/api")

	userGroup := api.Group("/users")
	userHttp.SetupUserRoutes(userGroup, userHandlers)

	health := api.Group("/healthz")
	health.GET("", healthHandler())

	return nil
}

// Health
//
//	@Summary		Check API health
//	@Description	Check if API is up
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/healthz [get]
func healthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}
