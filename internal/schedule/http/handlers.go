package http

import (
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/schedule"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type scheduleHandlers struct {
	cfg             *config.Config
	scheduleUseCase schedule.UseCase
	logger          logger.Logger
}

func NewScheduleHandlers(cfg *config.Config, scheduleUseCase schedule.UseCase, logger logger.Logger) schedule.Handlers {
	return &scheduleHandlers{cfg: cfg, scheduleUseCase: scheduleUseCase, logger: logger}
}

// Create
//
//	@Summary		Create new schedule
//	@Description	create new schedule
//	@Tags			Schedule
//	@Accept			json
//	@Produce		json
//	@Param			schedule	body		models.ScheduleCreate	true	"Schedule infos"
//	@Success		201			{object}	models.Schedule
//	@Failure		400			{object}	errorHandler.HttpErr
//	@Failure		500			{object}	errorHandler.HttpErr
//	@Router			/schedules [post]
func (u *scheduleHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var body models.ScheduleCreate

		scheduleCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		scheduleDb, err := u.scheduleUseCase.Create(user, &scheduleCreate)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, scheduleDb)
	}
}

// Unattended
//
//	@Summary		Get UnAttended Schedules
//	@Description	Get UnAttended Schedules
//	@Tags			Schedule
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	models.Schedule
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/schedules/ [get]
func (u *scheduleHandlers) GetUnattended() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		scheduleDb, err := u.scheduleUseCase.GetUnattended(user)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, scheduleDb)
	}
}

// Sign
//
//	@Summary		Sign for schedule
//	@Description	Sign for schedule
//	@Tags			Schedule
//	@Accept			json
//	@Produce		json
//	@Param			schedule	body		models.ScheduleCreate	true	"Schedule infos"
//	@Success		201			{object}	models.Schedule
//	@Failure		400			{object}	errorHandler.HttpErr
//	@Failure		500			{object}	errorHandler.HttpErr
//	@Router			/schedules/{id}/sign [post]
func (u *scheduleHandlers) Sign() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var body models.ScheduleSignatureCreate

		signature, err := request.ValidateJSON(body, ctx)
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

		scheduleDb, err := u.scheduleUseCase.Sign(&signature, user, uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, scheduleDb)
	}
}

// Check Sign
//
//	@Summary		Check Sign for schedule
//	@Description	Check Sign for schedule
//	@Tags			Schedule
//	@Accept			json
//	@Produce		json
//	@Param			schedule	body		models.ScheduleCreate	true	"Schedule infos"
//	@Success		201			{object}	models.Schedule
//	@Failure		400			{object}	errorHandler.HttpErr
//	@Failure		500			{object}	errorHandler.HttpErr
//	@Router			/schedules/{id}/sign [post]
func (u *scheduleHandlers) CheckSign() gin.HandlerFunc {
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

		scheduleDb, err := u.scheduleUseCase.CheckSign(user, uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, scheduleDb)
	}
}

// Read
//
//	@Summary		Get all schedule
//	@Description	Get all schedule
//	@Tags			Schedule
//	@Produce		json
//	@Success		200	{object}	[]models.Schedule
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/schedules [get]
func (u *scheduleHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		schedules, err := u.scheduleUseCase.GetAll(user)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, schedules)
	}
}

// Read
//
//	@Summary		Get schedule by id
//	@Description	Get schedule by id
//	@Tags			Schedule
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.Schedule
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/schedules/{id} [get]
func (u *scheduleHandlers) GetById() gin.HandlerFunc {
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

		schedule, err := u.scheduleUseCase.GetById(user, uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, schedule)
	}
}

// Read Code
//
//	@Summary		Get schedule by id
//	@Description	Get schedule by id
//	@Tags			Schedule
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.Schedule
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/schedules/{id}/code [get]
func (u *scheduleHandlers) GetSignatureCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.TEACHER)

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

		schedule, err := u.scheduleUseCase.GetSignatureCode(user, uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, schedule)
	}
}

// Update
//
//	@Summary		Update schedule
//	@Description	Update schedule
//	@Tags			Schedule
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int						true	"id"
//	@Param			schedule	body		models.ScheduleUpdate	true	"Schedule infos"
//	@Success		201			{object}	models.Schedule
//	@Failure		400			{object}	errorHandler.HttpErr
//	@Failure		500			{object}	errorHandler.HttpErr
//	@Router			/schedules/{id} [put]
func (u *scheduleHandlers) Update() gin.HandlerFunc {
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

		var body models.ScheduleUpdate

		scheduleUpdate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		schedule := &models.Schedule{
			Time:     *scheduleUpdate.Time,
			Duration: *scheduleUpdate.Duration,
			CourseId: *scheduleUpdate.CourseId,
			ClassId:  *scheduleUpdate.ClassId,
			CampusId: *scheduleUpdate.CampusId,
		}
		scheduleDb, err := u.scheduleUseCase.Update(user, uint(idInt), schedule)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, scheduleDb)
	}
}

// Delete
//
//	@Summary		Delete schedule by id
//	@Description	Delete schedule by id
//	@Tags			Schedule
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	nil
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/schedules/{id} [delete]
func (u *scheduleHandlers) Delete() gin.HandlerFunc {
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

		err = u.scheduleUseCase.Delete(user, uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
