package http

import (
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/class"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type classHandlers struct {
	cfg           *config.Config
	classUseCase  class.UseCase
	schoolUseCase school.UseCase
	logger        logger.Logger
}

func NewClassHandlers(cfg *config.Config, classUseCase class.UseCase, schoolUseCase school.UseCase, logger logger.Logger) class.Handlers {
	return &classHandlers{cfg: cfg, classUseCase: classUseCase, schoolUseCase: schoolUseCase, logger: logger}
}

// Create
//
//	@Summary		Create new class
//	@Description	create new class
//	@Tags			Class
//	@Accept			json
//	@Produce		json
//	@Param			class	body		models.ClassCreate	true	"Class infos"
//	@Success		201		{object}	models.Class
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/classes [post]
func (u *classHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var body models.ClassCreate

		classCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		school, err := u.schoolUseCase.GetByUser(user)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		class := &models.Class{
			Name:     classCreate.Name,
			PathId:   classCreate.PathId,
			SchoolId: school.ID,
		}
		classDb, err := u.classUseCase.Create(user, class)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, classDb)
	}
}

// Read
//
//	@Summary		Get all class
//	@Description	Get all class
//	@Tags			Class
//	@Produce		json
//	@Success		200	{object}	[]models.Class
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/classes [get]
func (u *classHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		school, err := u.schoolUseCase.GetByUser(user)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		class, err := u.classUseCase.GetAllBySchoolId(school.ID)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, class)
	}
}

// Read
//
//	@Summary		Get class by id
//	@Description	Get class by id
//	@Tags			Class
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.Class
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/classes/{id} [get]
func (u *classHandlers) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		school, err := u.schoolUseCase.GetByUser(user)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		class, err := u.classUseCase.GetById(uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		if school.ID != class.SchoolId {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			u.logger.Infof("Request: You can't retrieve class not in your school")
			return
		}

		ctx.JSON(http.StatusOK, class)
	}
}

// Add
//
//	@Summary		Add student to class
//	@Description	Add student to class
//	@Tags			Class
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"id"
//	@Param			class	body		models.ClassAdd	true	"Student infos"
//	@Success		201		{object}	models.Class
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/classes/{id}/add [post]
func (u *classHandlers) Add() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		// if user == nil || err != nil {
		// 	ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
		// 	return
		// }

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		var body models.ClassAdd

		classAdd, err := request.ValidateJSON(body, ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		classDb, err := u.classUseCase.Add(uint(idInt), &classAdd)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, classDb)
	}
}

// Update
//
//	@Summary		Update class
//	@Description	Update class
//	@Tags			Class
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"id"
//	@Param			class	body		models.ClassUpdate	true	"Class infos"
//	@Success		201		{object}	models.Class
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/classes/{id} [put]
func (u *classHandlers) Update() gin.HandlerFunc {
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

		school, err := u.schoolUseCase.GetByUser(user)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		classDb, err := u.classUseCase.GetById(uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		if classDb.SchoolId != school.ID {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			u.logger.Infof("Request: Not allowed to update class not on your school")
			return
		}

		var body models.ClassUpdate

		classUpdate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		class := &models.Class{
			Name:     classUpdate.Name,
			PathId:   classUpdate.PathId,
			SchoolId: classDb.SchoolId,
		}
		updatedClass, err := u.classUseCase.Update(uint(idInt), class)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, updatedClass)
	}
}

// Delete
//
//	@Summary		Delete class by id
//	@Description	Delete class by id
//	@Tags			Class
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	nil
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/classes/{id} [delete]
func (u *classHandlers) Delete() gin.HandlerFunc {
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

		err = u.classUseCase.Delete(user, uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
