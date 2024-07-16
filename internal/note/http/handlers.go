package http

import (
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/note"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type noteHandlers struct {
	cfg            *config.Config
	noteUseCase note.UseCase
	logger         logger.Logger
}

func NewNoteHandlers(cfg *config.Config, noteUseCase note.UseCase, logger logger.Logger) note.Handlers {
  return &noteHandlers{cfg: cfg, noteUseCase: noteUseCase, logger: logger}
}

// Create
//
//	@Summary		Create new note
//	@Description	create new note
//	@Tags			Note
//	@Accept			json
//	@Produce		json
//	@Param			note	body		models.NoteCreate	true	"Note infos"
//	@Success		201		{object}	models.Note
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/notes [post]
func (u *noteHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.TEACHER)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var body models.NoteCreate

		noteCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		note := &models.Note{
      Value: noteCreate.Value,
      TeacherId: user.ID,
			StudentId:  noteCreate.StudentId,
			ProjectId: noteCreate.ProjectId,
		}
		noteDb, err := u.noteUseCase.Create(note)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, noteDb)
	}
}

// Read
//
//	@Summary		Get all note
//	@Description	Get all note
//	@Tags			Note
//	@Produce		json
//	@Success		200	{object}	[]models.Note
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/notes [get]
func (u *noteHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		notes, err := u.noteUseCase.GetAllByUser(user)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, notes)
	}
}

// Update
//
//	@Summary		Update note
//	@Description	Update note
//	@Tags			Note
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Param			note	body		models.NoteUpdate	true	"Note infos"
//	@Success		201		{object}	models.Note
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/notes/{id} [put]
func (u *noteHandlers) Update() gin.HandlerFunc {
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

		var body models.NoteUpdate

		noteUpdate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		note := &models.Note{
      Value: noteUpdate.Value,
			ProjectId:       noteUpdate.ProjectId,
			StudentId: noteUpdate.StudentId,
      TeacherId: user.ID,
		}
		noteDb, err := u.noteUseCase.Update(uint(idInt), note)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, noteDb)
	}
}

// Delete
//
//	@Summary		Delete note by id
//	@Description	Delete note by id
//	@Tags			Note
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	nil
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/notes/{id} [delete]
func (u *noteHandlers) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		err = u.noteUseCase.Delete(uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
