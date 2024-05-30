package http

import (
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/path"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type pathHandlers struct {
	cfg         *config.Config
	pathUseCase path.UseCase
	logger      logger.Logger
}

func NewPathHandlers(cfg *config.Config, pathUseCase path.UseCase, logger logger.Logger) path.Handlers {
	return &pathHandlers{cfg: cfg, pathUseCase: pathUseCase, logger: logger}
}

// Create
//
//	@Summary		Create new path
//	@Description	create new path
//	@Tags			Path
//	@Accept			json
//	@Produce		json
//	@Param			path	body		models.PathCreate	true	"Path infos"
//	@Success		201		{object}	models.Path
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/paths [post]
func (u *pathHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var body models.PathCreate

		pathCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		path := &models.Path{
			Name:        pathCreate.Name,
			Description: pathCreate.Description,
			SchoolId:    pathCreate.SchoolId,
		}
		pathDb, err := u.pathUseCase.Create(user, path)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, pathDb)
	}
}

// Read
//
//	@Summary		Get all path
//	@Description	Get all path
//	@Tags			Path
//	@Produce		json
//	@Success		200	{object}	[]models.Path
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/paths [get]
func (u *pathHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paths, err := u.pathUseCase.GetAll()

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, paths)
	}
}

// Read
//
//	@Summary		Get path by id
//	@Description	Get path by id
//	@Tags			Path
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.Path
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/paths/{id} [get]
func (u *pathHandlers) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		path, err := u.pathUseCase.GetById(uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, path)
	}
}

// Update
//
//	@Summary		Update path
//	@Description	Update path
//	@Tags			Path
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"id"
//	@Param			path	body		models.PathUpdate	true	"Path infos"
//	@Success		201		{object}	models.Path
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/paths/{id} [put]
func (u *pathHandlers) Update() gin.HandlerFunc {
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

		var body models.PathUpdate

		pathUpdate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		path := &models.Path{
			Name:        pathUpdate.Name,
			Description: pathUpdate.Description,
			SchoolId:    pathUpdate.SchoolId,
		}
		pathDb, err := u.pathUseCase.Update(user, uint(idInt), path)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, pathDb)
	}
}

// Delete
//
//	@Summary		Delete path by id
//	@Description	Delete path by id
//	@Tags			Path
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	nil
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/paths/{id} [delete]
func (u *pathHandlers) Delete() gin.HandlerFunc {
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

		err = u.pathUseCase.Delete(user, uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
