package http

import (
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/example"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type exampleHandlers struct {
	exampleUseCase example.UseCase
	cfg            *config.Config
	logger         logger.Logger
}

func NewExampleHandlers(exampleUseCase example.UseCase, cfg *config.Config, logger logger.Logger) example.Handlers {
	return &exampleHandlers{exampleUseCase: exampleUseCase, cfg: cfg, logger: logger}
}

// Create
//
//	@Summary		Create new example
//	@Description	create new example
//	@Tags			Example
//	@Accept			json
//	@Produce		json
//	@Param			example	body		models.ExampleCreate	true	"Example infos"
//	@Success		201		{object}	models.Example
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/examples [post]
func (u *exampleHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.ExampleCreate

		exampleCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		example := &models.Example{
			Title:       exampleCreate.Title,
			Description: exampleCreate.Description,
		}
		exampleDb, err := u.exampleUseCase.Create(example)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, exampleDb)
	}
}

// Read
//
//	@Summary		Get all example
//	@Description	Get all example
//	@Tags			Example
//	@Produce		json
//	@Success		200	{object}	[]models.Example
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/examples [get]
func (u *exampleHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		examples, err := u.exampleUseCase.GetAll()

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, examples)
	}
}

// Read
//
//	@Summary		Get example by id
//	@Description	Get example by id
//	@Tags			Example
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.Example
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/examples/{id} [get]
func (u *exampleHandlers) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		example, err := u.exampleUseCase.GetById(uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, example)
	}
}

// Update
//
//	@Summary		Update example
//	@Description	Update example
//	@Tags			Example
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Param			example	body		models.ExampleUpdate	true	"Example infos"
//	@Success		201		{object}	models.Example
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/examples/{id} [put]
func (u *exampleHandlers) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		var body models.ExampleUpdate

		exampleUpdate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		example := &models.Example{
			Title:       exampleUpdate.Title,
			Description: exampleUpdate.Description,
		}
		exampleDb, err := u.exampleUseCase.Update(uint(idInt), example)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, exampleDb)
	}
}

// Delete
//
//	@Summary		Delete example by id
//	@Description	Delete example by id
//	@Tags			Example
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	nil
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/examples/{id} [delete]
func (u *exampleHandlers) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		err = u.exampleUseCase.Delete(uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
