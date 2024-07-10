package http

import (
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/chat"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type chatHandlers struct {
	cfg         *config.Config
	chatUseCase chat.UseCase
	logger      logger.Logger
}

func NewChatHandlers(cfg *config.Config, chatUseCase chat.UseCase, logger logger.Logger) chat.Handlers {
	return &chatHandlers{cfg: cfg, chatUseCase: chatUseCase, logger: logger}
}

// Create
//
//	@Summary		Create new channel
//	@Description	create new channel
//	@Tags			Chat
//	@Accept			json
//	@Produce		json
//	@Param			chat	body		models.ChannelCreate true	"Channel infos"
//	@Success		201		{object}	models.Channel
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/chats/channel [post]
func (u *chatHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var body models.ChannelCreate

		channelCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		channel := &models.Channel{
			FirstUserId:  channelCreate.FirstUserId,
			SecondUserId: channelCreate.SecondUserId,
		}
		channelDb, err := u.chatUseCase.Create(channel)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, channelDb)
	}
}

// Read
//
//	@Summary		Get all channels
//	@Description	Get all channels
//	@Tags			Chat
//	@Produce		json
//	@Success		200	{object}	[]models.Channel
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/chats/channel [get]
func (u *chatHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		channels, err := u.chatUseCase.GetAllByUser(user.ID)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, channels)
	}
}

// Read to chat
//
//	@Summary		Get all possible students chatter
//	@Description	Get all possible students chatter
//	@Tags			Chat
//	@Produce		json
//	@Success		200	{object}	[]models.User
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/chats/students [get]
func (u *chatHandlers) GetAllStudentChatter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.TEACHER)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		students, err := u.chatUseCase.GetAllPossibleChatStudent(user)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, students)
	}
}

// Read to chat
//
//	@Summary		Get all possible teacher chatter
//	@Description	Get all possible teacher chatter
//	@Tags			Chat
//	@Produce		json
//	@Success		200	{object}	[]models.User
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/chats/teachers [get]
func (u *chatHandlers) GetAllTeacherChatter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		teachers, err := u.chatUseCase.GetAllPossibleChatTeacher(user)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, teachers)
	}
}

// Read
//
//	@Summary		Get channel by id
//	@Description	Get channel by id
//	@Tags			Chat
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.Channel
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/chats/channel/{id} [get]
func (u *chatHandlers) GetById() gin.HandlerFunc {
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

		channel, err := u.chatUseCase.GetById(uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, channel)
	}
}
