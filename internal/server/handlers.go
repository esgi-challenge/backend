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
	chatHttp "github.com/esgi-challenge/backend/internal/chat/http"
	chatRepo "github.com/esgi-challenge/backend/internal/chat/repository"
	chatUseCase "github.com/esgi-challenge/backend/internal/chat/usecase"
	classHttp "github.com/esgi-challenge/backend/internal/class/http"
	classRepo "github.com/esgi-challenge/backend/internal/class/repository"
	classUseCase "github.com/esgi-challenge/backend/internal/class/usecase"
	courseHttp "github.com/esgi-challenge/backend/internal/course/http"
	courseRepo "github.com/esgi-challenge/backend/internal/course/repository"
	courseUseCase "github.com/esgi-challenge/backend/internal/course/usecase"
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

	informationsHttp "github.com/esgi-challenge/backend/internal/informations/http"
	informationsRepo "github.com/esgi-challenge/backend/internal/informations/repository"
	informationsUseCase "github.com/esgi-challenge/backend/internal/informations/usecase"

	projectHttp "github.com/esgi-challenge/backend/internal/project/http"
	projectRepo "github.com/esgi-challenge/backend/internal/project/repository"
	projectUseCase "github.com/esgi-challenge/backend/internal/project/usecase"

	documentHttp "github.com/esgi-challenge/backend/internal/document/http"
	documentRepo "github.com/esgi-challenge/backend/internal/document/repository"
	documentUseCase "github.com/esgi-challenge/backend/internal/document/usecase"

	noteHttp "github.com/esgi-challenge/backend/internal/note/http"
	noteRepo "github.com/esgi-challenge/backend/internal/note/repository"
	noteUseCase "github.com/esgi-challenge/backend/internal/note/usecase"

	"github.com/esgi-challenge/backend/internal/websocket"
)

func (s *Server) SetupHandlers() error {
	// Repo
	userRepo := userRepo.NewUserRepository(s.psqlDB)
	schoolRepo := schoolRepo.NewSchoolRepository(s.psqlDB)
	campusRepo := campusRepo.NewCampusRepository(s.psqlDB)
	pathRepo := pathRepo.NewPathRepository(s.psqlDB)
	classRepo := classRepo.NewClassRepository(s.psqlDB)
	courseRepo := courseRepo.NewCourseRepository(s.psqlDB)
	scheduleRepo := scheduleRepo.NewScheduleRepository(s.psqlDB)
	informationsRepo := informationsRepo.NewInformationsRepository(s.psqlDB)
	chatRepo := chatRepo.NewChatRepository(s.psqlDB)
	projectRepo := projectRepo.NewProjectRepository(s.psqlDB)
	documentRepo := documentRepo.NewDocumentRepository(s.psqlDB)
	noteRepo := noteRepo.NewNoteRepository(s.psqlDB)

	// UseCase
	userUseCase := userUseCase.NewUserUseCase(userRepo, s.cfg, s.logger)
	schoolUseCase := schoolUseCase.NewSchoolUseCase(s.cfg, schoolRepo, userRepo, s.logger)
	authUseCase := authUseCase.NewAuthUseCase(s.cfg, userRepo, s.logger)
	campusUseCase := campusUseCase.NewCampusUseCase(s.cfg, campusRepo, schoolRepo, s.logger)
	pathUseCase := pathUseCase.NewPathUseCase(s.cfg, pathRepo, schoolRepo, s.logger)
	classUseCase := classUseCase.NewClassUseCase(s.cfg, classRepo, pathRepo, schoolRepo, userRepo, s.logger)
	courseUseCase := courseUseCase.NewCourseUseCase(s.cfg, courseRepo, pathRepo, schoolRepo, s.logger)
	scheduleUseCase := scheduleUseCase.NewScheduleUseCase(s.cfg, scheduleRepo, courseRepo, pathRepo, schoolRepo, campusRepo, userRepo, s.logger)
	informationsUseCase := informationsUseCase.NewInformationsUseCase(s.cfg, informationsRepo, schoolRepo, s.logger)
	chatUseCase := chatUseCase.NewChatUseCase(s.cfg, chatRepo, schoolRepo, s.logger)
	documentUseCase := documentUseCase.NewDocumentUseCase(s.cfg, documentRepo, courseRepo, schoolRepo, s.logger, *s.storage)
	projectsUseCase := projectUseCase.NewProjectUseCase(s.cfg, projectRepo, courseUseCase, classUseCase, documentUseCase, s.logger)
	noteUseCase := noteUseCase.NewNoteUseCase(s.cfg, noteRepo, s.logger)

	// Handlers
	userHandlers := userHttp.NewUserHandlers(userUseCase, s.cfg, s.logger)
	schoolHandlers := schoolHttp.NewSchoolHandlers(s.cfg, schoolUseCase, userUseCase, s.logger)
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUseCase, s.logger)
	campusHandlers := campusHttp.NewCampusHandlers(s.cfg, campusUseCase, schoolUseCase, s.logger, s.gmapApiManager)
	pathHandlers := pathHttp.NewPathHandlers(s.cfg, pathUseCase, schoolUseCase, s.logger)
	classHandlers := classHttp.NewClassHandlers(s.cfg, classUseCase, schoolUseCase, s.logger)
	courseHandlers := courseHttp.NewCourseHandlers(s.cfg, courseUseCase, schoolUseCase, s.logger)
	scheduleHandlers := scheduleHttp.NewScheduleHandlers(s.cfg, scheduleUseCase, schoolUseCase, s.logger)
	informationsHandlers := informationsHttp.NewInformationsHandlers(s.cfg, informationsUseCase, schoolUseCase, s.logger)
	chatHandlers := chatHttp.NewChatHandlers(s.cfg, chatUseCase, s.logger)
	projectHandlers := projectHttp.NewProjectHandlers(s.cfg, projectsUseCase, s.logger)
	documentHandlers := documentHttp.NewDocumentHandlers(s.cfg, documentUseCase, s.logger)
	noteHandler := noteHttp.NewNoteHandlers(s.cfg, noteUseCase, s.logger)

	// Middlewares
	mw := middleware.InitMiddlewareManager(s.cfg, s.logger)

	s.router.Use(mw.RequestMiddleware())
	s.router.Use(mw.CorsMiddleware())
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := s.router.Group("/api")
	wellknown := s.router.Group("/.well-known")

	userGroup := api.Group("/users")
	schoolGroup := api.Group("/schools")
	authGroup := api.Group("/auth")
	campusGroup := api.Group("/campus")
	pathGroup := api.Group("/paths")
	classGroup := api.Group("/classes")
	courseGroup := api.Group("/courses")
	schedulesGroup := api.Group("/schedules")
	informationsGroup := api.Group("/informations")
	chatGroup := api.Group("/chats")
	projectGroup := api.Group("/projects")
	documentGroup := api.Group("/documents")
	noteGroup := api.Group("/notes")

	websocketHandlers := &websocket.WebSocketHandler{
		Cfg:         s.cfg,
		ChatUseCase: chatUseCase,
		Logger:      s.logger,
	}
	websocketGroup := api.Group("/ws/chat")
	websocketGroup.GET("/:channelId", websocketHandlers.ChatHandler)

	userHttp.SetupUserRoutes(userGroup, userHandlers)
	schoolHttp.SetupSchoolRoutes(schoolGroup, schoolHandlers)
	authHttp.SetupAuthRoutes(authGroup, authHandlers)
	campusHttp.SetupCampusRoutes(campusGroup, campusHandlers)
	pathHttp.SetupPathRoutes(pathGroup, pathHandlers)
	classHttp.SetupClassRoutes(classGroup, classHandlers)
	courseHttp.SetupCourseRoutes(courseGroup, courseHandlers)
	scheduleHttp.SetupScheduleRoutes(schedulesGroup, scheduleHandlers)
	informationsHttp.SetupInformationsRoutes(informationsGroup, informationsHandlers)
	chatHttp.SetupChatRoutes(chatGroup, chatHandlers)
	projectHttp.SetupProjectRoutes(projectGroup, projectHandlers)
	documentHttp.SetupDocumentRoutes(documentGroup, documentHandlers)
	noteHttp.SetupNoteRoutes(noteGroup, noteHandler)

	wk.SetupPathRoutes(wellknown)

	health := api.Group("/healthz")
	health.GET("", healthHandler())

	s.logger.Info("Checking if admin existing...")
	_, err := userRepo.GetByEmail(s.cfg.AdminEmail)

	var userkind models.UserKind = models.SUPERADMIN

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.Infof("There is no admin user, creating one...")
		admin := &models.User{
			Firstname: "admin",
			Lastname:  "admin",
			Email:     s.cfg.AdminEmail,
			Password:  s.cfg.AdminPassword,
			UserKind:  &userkind,
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
