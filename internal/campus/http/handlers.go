package http

import (
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/campus"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/gin-gonic/gin"
)

type campusHandlers struct {
	cfg           *config.Config
	campusUseCase campus.UseCase
  schoolUseCase school.UseCase
	logger        logger.Logger
}

func NewCampusHandlers(cfg *config.Config, campusUseCase campus.UseCase, schoolUseCase school.UseCase, logger logger.Logger) campus.Handlers {
	return &campusHandlers{cfg: cfg, campusUseCase: campusUseCase, schoolUseCase: schoolUseCase, logger: logger}
}

// Create
//
//	@Summary		Create new campus
//	@Description	create new campus
//	@Tags			Campus
//	@Accept			json
//	@Produce		json
//	@Param			campus	body		models.CampusCreate	true	"Campus infos"
//	@Success		201		{object}	models.Campus
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/campus [post]
func (u *campusHandlers) Create() gin.HandlerFunc {
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

		var body models.CampusCreate

		campusCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		campus := &models.Campus{
			Name:     campusCreate.Name,
			Location: campusCreate.Location,
			SchoolId: school.ID,
      Latitude: campusCreate.Latitude,
      Longitude: campusCreate.Longitude,
		}

		campusDb, err := u.campusUseCase.Create(user, campus)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, campusDb)
	}
}

// Read
//
//	@Summary		Get all campus
//	@Description	Get all campus
//	@Tags			Campus
//	@Produce		json
//	@Success		200	{object}	[]models.Campus
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/campus [get]
func (u *campusHandlers) GetAll() gin.HandlerFunc {
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
		campus, err := u.campusUseCase.GetAllBySchoolId(school.ID)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, campus)
	}
}

// Read
//
//	@Summary		Get campus by id
//	@Description	Get campus by id
//	@Tags			Campus
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.Campus
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/campus/{id} [get]
func (u *campusHandlers) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		campus, err := u.campusUseCase.GetById(uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, campus)
	}
}

// Update
//
//	@Summary		Update campus
//	@Description	Update campus
//	@Tags			Campus
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"id"
//	@Param			campus	body		models.CampusUpdate	true	"Campus infos"
//	@Success		201		{object}	models.Campus
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/campus/{id} [put]
func (u *campusHandlers) Update() gin.HandlerFunc {
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

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		var body models.CampusUpdate

		campusUpdate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		campusDb, err := u.campusUseCase.GetById(uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		if campusDb.SchoolId != school.ID {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			u.logger.Infof("Request: Not allowed to update campus not on your school")
			return
		}

		campus := &models.Campus{
			Name:     campusUpdate.Name,
			Location: campusUpdate.Location,
      Latitude: campusUpdate.Latitude,
      Longitude: campusUpdate.Longitude,
			SchoolId: campusDb.SchoolId,
		}
		updatedCampus, err := u.campusUseCase.Update(user, uint(idInt), campus)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, updatedCampus)
	}
}

// Delete
//
//	@Summary		Delete campus by id
//	@Description	Delete campus by id
//	@Tags			Campus
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	nil
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/campus/{id} [delete]
func (u *campusHandlers) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		school, err := u.schoolUseCase.GetByUser(user)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		campusDb, err := u.campusUseCase.GetById(uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		if campusDb.SchoolId != school.ID {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			u.logger.Infof("Request: Not allowed to delete campus not on your school")
			return
		}

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		err = u.campusUseCase.Delete(user, uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
