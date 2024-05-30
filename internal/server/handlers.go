package server

import (
	"net/http"

	_ "github.com/esgi-challenge/backend/docs"
	"github.com/esgi-challenge/backend/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	authHttp "github.com/esgi-challenge/backend/internal/auth/http"
	authUseCase "github.com/esgi-challenge/backend/internal/auth/usecase"
	campusHttp "github.com/esgi-challenge/backend/internal/campus/http"
	campusRepo "github.com/esgi-challenge/backend/internal/campus/repository"
	campusUseCase "github.com/esgi-challenge/backend/internal/campus/usecase"
	classHttp "github.com/esgi-challenge/backend/internal/class/http"
	classRepo "github.com/esgi-challenge/backend/internal/class/repository"
	classUseCase "github.com/esgi-challenge/backend/internal/class/usecase"
	exampleHttp "github.com/esgi-challenge/backend/internal/example/http"
	exampleRepo "github.com/esgi-challenge/backend/internal/example/repository"
	exampleUseCase "github.com/esgi-challenge/backend/internal/example/usecase"
	pathHttp "github.com/esgi-challenge/backend/internal/path/http"
	pathRepo "github.com/esgi-challenge/backend/internal/path/repository"
	pathUseCase "github.com/esgi-challenge/backend/internal/path/usecase"
	schoolHttp "github.com/esgi-challenge/backend/internal/school/http"
	schoolRepo "github.com/esgi-challenge/backend/internal/school/repository"
	schoolUseCase "github.com/esgi-challenge/backend/internal/school/usecase"
	userHttp "github.com/esgi-challenge/backend/internal/user/http"
	userRepo "github.com/esgi-challenge/backend/internal/user/repository"
	userUseCase "github.com/esgi-challenge/backend/internal/user/usecase"
)

func (s *Server) SetupHandlers() error {
	// Repo
	exampleRepo := exampleRepo.NewExampleRepository(s.psqlDB)
	userRepo := userRepo.NewUserRepository(s.psqlDB)
	schoolRepo := schoolRepo.NewSchoolRepository(s.psqlDB)
	campusRepo := campusRepo.NewCampusRepository(s.psqlDB)
	pathRepo := pathRepo.NewPathRepository(s.psqlDB)
	classRepo := classRepo.NewClassRepository(s.psqlDB)

	// UseCase
	exampleUseCase := exampleUseCase.NewExampleUseCase(s.cfg, exampleRepo, s.logger)
	userUseCase := userUseCase.NewUserUseCase(userRepo, s.cfg, s.logger)
	schoolUseCase := schoolUseCase.NewSchoolUseCase(s.cfg, schoolRepo, userRepo, s.logger)
	authUseCase := authUseCase.NewAuthUseCase(s.cfg, userRepo, s.logger)
	campusUseCase := campusUseCase.NewCampusUseCase(s.cfg, campusRepo, schoolRepo, s.logger)
	pathUseCase := pathUseCase.NewPathUseCase(s.cfg, pathRepo, schoolRepo, s.logger)
	classUseCase := classUseCase.NewClassUseCase(s.cfg, classRepo, pathRepo, schoolRepo, userRepo, s.logger)

	// Handlers
	exampleHandlers := exampleHttp.NewExampleHandlers(s.cfg, exampleUseCase, s.logger)
	userHandlers := userHttp.NewUserHandlers(userUseCase, s.cfg, s.logger)
	schoolHandlers := schoolHttp.NewSchoolHandlers(s.cfg, schoolUseCase, s.logger)
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUseCase, s.logger)
	campusHandlers := campusHttp.NewCampusHandlers(s.cfg, campusUseCase, s.logger)
	pathHandlers := pathHttp.NewPathHandlers(s.cfg, pathUseCase, s.logger)
	classHandlers := classHttp.NewClassHandlers(s.cfg, classUseCase, s.logger)

	mw := middleware.InitMiddlewareManager(s.cfg, s.logger)

	s.router.Use(mw.RequestMiddleware())
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := s.router.Group("/api")

	exampleGroup := api.Group("/examples")
	userGroup := api.Group("/users")
	schoolGroup := api.Group("/schools")
	authGroup := api.Group("/auth")
	campusGroup := api.Group("/campus")
	pathGroup := api.Group("/path")
	classGroup := api.Group("/class")

	exampleHttp.SetupExampleRoutes(exampleGroup, exampleHandlers)
	userHttp.SetupUserRoutes(userGroup, userHandlers)
	schoolHttp.SetupSchoolRoutes(schoolGroup, schoolHandlers)
	authHttp.SetupAuthRoutes(authGroup, authHandlers)
	campusHttp.SetupCampusRoutes(campusGroup, campusHandlers)
	pathHttp.SetupPathRoutes(pathGroup, pathHandlers)
	classHttp.SetupClassRoutes(classGroup, classHandlers)

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
