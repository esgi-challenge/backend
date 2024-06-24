package server

import (
	"errors"
	"net/http"

	_ "github.com/esgi-challenge/backend/docs"
	"github.com/esgi-challenge/backend/internal/middleware"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	authHttp "github.com/esgi-challenge/backend/internal/auth/http"
	authUseCase "github.com/esgi-challenge/backend/internal/auth/usecase"
	campusHttp "github.com/esgi-challenge/backend/internal/campus/http"
	campusRepo "github.com/esgi-challenge/backend/internal/campus/repository"
	campusUseCase "github.com/esgi-challenge/backend/internal/campus/usecase"
	classHttp "github.com/esgi-challenge/backend/internal/class/http"
	classRepo "github.com/esgi-challenge/backend/internal/class/repository"
	classUseCase "github.com/esgi-challenge/backend/internal/class/usecase"
	courseHttp "github.com/esgi-challenge/backend/internal/course/http"
	courseRepo "github.com/esgi-challenge/backend/internal/course/repository"
	courseUseCase "github.com/esgi-challenge/backend/internal/course/usecase"
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
	wk "github.com/esgi-challenge/backend/internal/well-known"

	scheduleHttp "github.com/esgi-challenge/backend/internal/schedule/http"
	scheduleRepo "github.com/esgi-challenge/backend/internal/schedule/repository"
	scheduleUseCase "github.com/esgi-challenge/backend/internal/schedule/usecase"
)

func (s *Server) SetupHandlers() error {
	// Repo
	exampleRepo := exampleRepo.NewExampleRepository(s.psqlDB)
	userRepo := userRepo.NewUserRepository(s.psqlDB)
	schoolRepo := schoolRepo.NewSchoolRepository(s.psqlDB)
	campusRepo := campusRepo.NewCampusRepository(s.psqlDB)
	pathRepo := pathRepo.NewPathRepository(s.psqlDB)
	classRepo := classRepo.NewClassRepository(s.psqlDB)
	courseRepo := courseRepo.NewCourseRepository(s.psqlDB)
	scheduleRepo := scheduleRepo.NewScheduleRepository(s.psqlDB)

	// UseCase
	exampleUseCase := exampleUseCase.NewExampleUseCase(s.cfg, exampleRepo, s.logger)
	userUseCase := userUseCase.NewUserUseCase(userRepo, s.cfg, s.logger)
	schoolUseCase := schoolUseCase.NewSchoolUseCase(s.cfg, schoolRepo, userRepo, s.logger)
	authUseCase := authUseCase.NewAuthUseCase(s.cfg, userRepo, s.logger)
	campusUseCase := campusUseCase.NewCampusUseCase(s.cfg, campusRepo, schoolRepo, s.logger)
	pathUseCase := pathUseCase.NewPathUseCase(s.cfg, pathRepo, schoolRepo, s.logger)
	classUseCase := classUseCase.NewClassUseCase(s.cfg, classRepo, pathRepo, schoolRepo, userRepo, s.logger)
	courseUseCase := courseUseCase.NewCourseUseCase(s.cfg, courseRepo, pathRepo, schoolRepo, s.logger)
	scheduleUseCase := scheduleUseCase.NewScheduleUseCase(s.cfg, scheduleRepo, courseRepo, pathRepo, schoolRepo, s.logger)

	// Handlers
	exampleHandlers := exampleHttp.NewExampleHandlers(s.cfg, exampleUseCase, s.logger)
	userHandlers := userHttp.NewUserHandlers(userUseCase, s.cfg, s.logger)
	schoolHandlers := schoolHttp.NewSchoolHandlers(s.cfg, schoolUseCase, s.logger)
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUseCase, s.logger)
	campusHandlers := campusHttp.NewCampusHandlers(s.cfg, campusUseCase, s.logger)
	pathHandlers := pathHttp.NewPathHandlers(s.cfg, pathUseCase, s.logger)
	classHandlers := classHttp.NewClassHandlers(s.cfg, classUseCase, s.logger)
	courseHandlers := courseHttp.NewCourseHandlers(s.cfg, courseUseCase, s.logger)
	scheduleHandlers := scheduleHttp.NewScheduleHandlers(s.cfg, scheduleUseCase, s.logger)

	// Middlewares
	mw := middleware.InitMiddlewareManager(s.cfg, s.logger)

	s.router.Use(mw.RequestMiddleware())
	s.router.Use(mw.CorsMiddleware())
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := s.router.Group("/api")
	wellknown := s.router.Group("/.well-known")

	exampleGroup := api.Group("/examples")
	userGroup := api.Group("/users")
	schoolGroup := api.Group("/schools")
	authGroup := api.Group("/auth")
	campusGroup := api.Group("/campus")
	pathGroup := api.Group("/path")
	classGroup := api.Group("/classes")
	courseGroup := api.Group("/courses")
	schedulesGroup := api.Group("/schedules")

	exampleHttp.SetupExampleRoutes(exampleGroup, exampleHandlers)
	userHttp.SetupUserRoutes(userGroup, userHandlers)
	schoolHttp.SetupSchoolRoutes(schoolGroup, schoolHandlers)
	authHttp.SetupAuthRoutes(authGroup, authHandlers)
	campusHttp.SetupCampusRoutes(campusGroup, campusHandlers)
	pathHttp.SetupPathRoutes(pathGroup, pathHandlers)
	classHttp.SetupClassRoutes(classGroup, classHandlers)
	courseHttp.SetupCourseRoutes(courseGroup, courseHandlers)
	scheduleHttp.SetupScheduleRoutes(schedulesGroup, scheduleHandlers)
	wk.SetupPathRoutes(wellknown)

	health := api.Group("/healthz")
	health.GET("", healthHandler())

	s.logger.Info("Checking if admin existing...")
	_, err := userRepo.GetByEmail(s.cfg.AdminEmail)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.Infof("There is no admin user, creating one...")
		admin := &models.User{
			Firstname: "admin",
			Lastname:  "admin",
			Email:     s.cfg.AdminEmail,
			Password:  s.cfg.AdminPassword,
			UserKind:  models.SUPERADMIN,
		}
		err := admin.HashPassword()
		if err != nil {
			s.logger.Fatalf("Admin error: %v", err)
		}

		_, err = userRepo.Create(admin)
		if err != nil {
			s.logger.Fatalf("Admin error: %v", err)
		}

		s.logger.Info("Admin user has been created !")
	} else if err != nil {
		s.logger.Fatalf("Admin error: %v", err)
	} else {
		s.logger.Info("Admin user already exist !")
	}

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
