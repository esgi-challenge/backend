package http

import (
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/informations"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type informationsHandlers struct {
	cfg                 *config.Config
	informationsUseCase informations.UseCase
	logger              logger.Logger
}

func NewInformationsHandlers(cfg *config.Config, informationsUseCase informations.UseCase, logger logger.Logger) informations.Handlers {
	return &informationsHandlers{cfg: cfg, informationsUseCase: informationsUseCase, logger: logger}
}

// Create
//
//	@Summary		Create new informations
//	@Description	create new informations
//	@Tags			Informations
//	@Accept			json
//	@Produce		json
//	@Param			informations	body		models.InformationsCreate	true	"Informations infos"
//	@Success		201		{object}	models.Informations
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/informationss [post]
func (u *informationsHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var body models.InformationsCreate

		informationsCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		informations := &models.Informations{
			Title:       informationsCreate.Title,
			Description: informationsCreate.Description,
			SchoolId:    informationsCreate.SchoolId,
		}

		informationsDb, err := u.informationsUseCase.Create(user, informations)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, informationsDb)
	}
}

// Read
//
//	@Summary		Get all informations
//	@Description	Get all informations
//	@Tags			Informations
//	@Produce		json
//	@Success		200	{object}	[]models.Informations
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/informationss [get]
func (u *informationsHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		informationss, err := u.informationsUseCase.GetAll(user)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, informationss)
	}
}

// Read
//
//	@Summary		Get informations by id
//	@Description	Get informations by id
//	@Tags			Informations
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.Informations
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/informationss/{id} [get]
func (u *informationsHandlers) GetById() gin.HandlerFunc {
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

		informations, err := u.informationsUseCase.GetById(user, uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, informations)
	}
}

// Update
//
//	@Summary		Update informations
//	@Description	Update informations
//	@Tags			Informations
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Param			informations	body		models.InformationsUpdate	true	"Informations infos"
//	@Success		201		{object}	models.Informations
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/informationss/{id} [put]
func (u *informationsHandlers) Update() gin.HandlerFunc {
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

		var body models.InformationsUpdate

		informationsUpdate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		informations := &models.Informations{
			Title:       informationsUpdate.Title,
			Description: informationsUpdate.Description,
			SchoolId:    informationsUpdate.SchoolId,
		}
		informationsDb, err := u.informationsUseCase.Update(user, uint(idInt), informations)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, informationsDb)
	}
}

// Delete
//
//	@Summary		Delete informations by id
//	@Description	Delete informations by id
//	@Tags			Informations
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	nil
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/informationss/{id} [delete]
func (u *informationsHandlers) Delete() gin.HandlerFunc {
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

		err = u.informationsUseCase.Delete(user, uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
