package server

import (
	"net/http"

	_ "github.com/esgi-challenge/backend/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	exampleHttp "github.com/esgi-challenge/backend/internal/example/http"
	exampleRepo "github.com/esgi-challenge/backend/internal/example/repository"
	exampleUseCase "github.com/esgi-challenge/backend/internal/example/usecase"
	"github.com/esgi-challenge/backend/internal/middleware"
	userHttp "github.com/esgi-challenge/backend/internal/user/http"
	userRepo "github.com/esgi-challenge/backend/internal/user/repository"
	userUseCase "github.com/esgi-challenge/backend/internal/user/usecase"
)

func (s *Server) SetupHandlers() error {
	// Repo
	exampleRepo := exampleRepo.NewExampleRepository(s.psqlDB)
	userRepo := userRepo.NewUserRepository(s.psqlDB)

	// UseCase
	exampleUseCase := exampleUseCase.NewExampleUseCase(exampleRepo, s.cfg, s.logger)
	userUseCase := userUseCase.NewUserUseCase(userRepo, s.cfg, s.logger)

	// Handlers
	exampleHandlers := exampleHttp.NewExampleHandlers(exampleUseCase, s.cfg, s.logger)
	userHandlers := userHttp.NewUserHandlers(userUseCase, s.cfg, s.logger)

	mw := middleware.InitMiddlewareManager(s.cfg, s.logger)

	s.router.Use(mw.RequestMiddleware())
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := s.router.Group("/api")

	exampleGroup := api.Group("/examples")
	userGroup := api.Group("/users")

	exampleHttp.SetupExampleRoutes(exampleGroup, exampleHandlers)
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
