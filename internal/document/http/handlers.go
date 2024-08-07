package http

import (
	"bufio"
	"io"
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/document"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type documentHandlers struct {
	cfg             *config.Config
	documentUseCase document.UseCase
	logger          logger.Logger
}

func NewDocumentHandlers(cfg *config.Config, documentUseCase document.UseCase, logger logger.Logger) document.Handlers {
	return &documentHandlers{
		cfg:             cfg,
		documentUseCase: documentUseCase,
		logger:          logger,
	}
}

// Create
//
//	@Summary		Create new document
//	@Description	create new document
//	@Tags			Document
//	@Accept			json
//	@Produce		json
//	@Param			document	body		models.DocumentCreate	true	"Document infos"
//	@Success		201			{object}	models.Document
//	@Failure		400			{object}	errorHandler.HttpErr
//	@Failure		500			{object}	errorHandler.HttpErr
//	@Router			/documents [post]
func (u *documentHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var courseIdUint *uint

		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		name := ctx.Request.FormValue("name")

		file, header, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		courseIdStr := ctx.Request.FormValue("courseId")

		if courseIdStr != "" {
			courseId, err := strconv.ParseUint(courseIdStr, 10, 32)

			if err != nil && err != io.EOF {
				ctx.AbortWithStatusJSON(errorHandler.InternalServerErrorResponse())
				u.logger.Infof("Request: %v", err.Error())
				return
			}
			tmp := uint(courseId)

			courseIdUint = &tmp
		}

		bs := make([]byte, header.Size)
		_, err = bufio.NewReader(file).Read(bs)

		if err != nil && err != io.EOF {
			ctx.AbortWithStatusJSON(errorHandler.InternalServerErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		documentDb, err := u.documentUseCase.Create(user, &models.DocumentCreate{
			Name:     name,
			Byte:     bs,
			CourseId: courseIdUint,
		})

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, documentDb)
	}
}

// Read All
//
//	@Summary		Get all documents
//	@Description	Get all documents
//	@Tags			Document
//	@Produce		json
//	@Success		200	{object}	[]models.Document
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/documents [get]
func (u *documentHandlers) GetAllByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.TEACHER)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		document, err := u.documentUseCase.GetAllByUserId(uint(user.ID))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, document)
	}
}

// Read All
//
//	@Summary		Get all documents
//	@Description	Get all documents
//	@Tags			Document
//	@Produce		json
//	@Success		200	{object}	[]models.Document
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/documents [get]
func (u *documentHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		documents, err := u.documentUseCase.GetAll(user)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, documents)
	}
}

// Read
//
//	@Summary		Get document by id
//	@Description	Get document by id
//	@Tags			Document
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.Document
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/documents/{id} [get]
func (u *documentHandlers) GetById() gin.HandlerFunc {
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

		document, err := u.documentUseCase.GetById(user, uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, document)
	}
}

// Delete
//
//	@Summary		Delete document by id
//	@Description	Delete document by id
//	@Tags			Document
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	nil
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/documents/{id} [delete]
func (u *documentHandlers) Delete() gin.HandlerFunc {
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

		err = u.documentUseCase.Delete(user, uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
