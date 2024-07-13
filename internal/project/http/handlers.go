package http

import (
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/project"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type projectHandlers struct {
	cfg            *config.Config
	projectUseCase project.UseCase
	logger         logger.Logger
}

func NewProjectHandlers(cfg *config.Config, projectUseCase project.UseCase, logger logger.Logger) project.Handlers {
	return &projectHandlers{cfg: cfg, projectUseCase: projectUseCase, logger: logger}
}

// Create
//
//	@Summary		Create new project
//	@Description	create new project
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			project	body		models.ProjectCreate	true	"Project infos"
//	@Success		201		{object}	models.Project
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/projects [post]
func (u *projectHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var body models.ProjectCreate

		projectCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		project := &models.Project{
			Title:      projectCreate.Title,
			CourseId:   *projectCreate.CourseId,
			ClassId:    *projectCreate.ClassId,
			DocumentId: *projectCreate.DocumentId,
			EndDate:    projectCreate.EndDate,
		}
		projectDb, err := u.projectUseCase.Create(user, project)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, projectDb)
	}
}

// Read
//
//	@Summary		Get all project
//	@Description	Get all project
//	@Tags			Project
//	@Produce		json
//	@Success		200	{object}	[]models.Project
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/projects [get]
func (u *projectHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		projects, err := u.projectUseCase.GetAll(user)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, projects)
	}
}

// Read
//
//	@Summary		Get project by id
//	@Description	Get project by id
//	@Tags			Project
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.Project
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/projects/{id} [get]
func (u *projectHandlers) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		project, err := u.projectUseCase.GetById(user, uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, project)
	}
}

// Get Groups
//
//	@Summary		Get project's groups
//	@Description	Get project's groups
//	@Tags			Project
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	[]models.ProjectGroup
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/projects/{id}/groups [get]
func (u *projectHandlers) GetGroups() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		project, err := u.projectUseCase.GetGroups(user, uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, project)
	}
}

// Join
//
//	@Summary		Join project
//	@Description	Join project
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			project	body		models.ProjectStudentCreate	true	"Join infos"
//	@Success		201		{object}	models.ProjectStudent
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/projects/{id}/join [post]
func (u *projectHandlers) JoinProject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var body models.ProjectStudentCreate

		projectStudentCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		join := &models.ProjectStudentCreate{
			Group: projectStudentCreate.Group,
		}
		projectDb, err := u.projectUseCase.JoinProject(user, join, uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, projectDb)
	}
}

// Quit
//
//	@Summary		Quit project
//	@Description	Quit project
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			project	body		models.ProjectStudentCreate	true	"Join infos"
//	@Success		200		{object}	nil
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		404		{object}	errorHandler.HttpErr
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Router			/projects/{id}/quit [post]
func (u *projectHandlers) QuitProject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		err = u.projectUseCase.QuitProject(user, uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}

// Update
//
//	@Summary		Update project
//	@Description	Update project
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		201	{object}	models.Project
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/projects/{id} [put]
func (u *projectHandlers) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		var body models.ProjectUpdate

		projectUpdate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		project := &models.Project{
			Title:      projectUpdate.Title,
			CourseId:   *projectUpdate.CourseId,
			ClassId:    *projectUpdate.ClassId,
			DocumentId: *projectUpdate.DocumentId,
			EndDate:    projectUpdate.EndDate,
		}
		projectDb, err := u.projectUseCase.Update(user, uint(idInt), project)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, projectDb)
	}
}

// Delete
//
//	@Summary		Delete project by id
//	@Description	Delete project by id
//	@Tags			Project
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	nil
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/projects/{id} [delete]
func (u *projectHandlers) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		err = u.projectUseCase.Delete(user, uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
