package http

import (
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/esgi-challenge/backend/pkg/email"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type userHandlers struct {
	userUseCase user.UseCase
	cfg         *config.Config
	logger      logger.Logger
}

func NewUserHandlers(userUseCase user.UseCase, cfg *config.Config, logger logger.Logger) user.Handlers {
	return &userHandlers{userUseCase: userUseCase, cfg: cfg, logger: logger}
}

// Read
//
//	@Summary		Get all users
//	@Description	Get all users
//	@Tags			User
//	@Produce		json
//	@Success		200	{object}	[]models.User
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/users [get]
func (u *userHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := u.userUseCase.GetAll()

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, users)
	}
}

// Create
//
//	@Summary		Create user
//	@Description	Create User
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user body		models.UserCreate	true	"User Infos"
//	@Success		201		{object}	models.User
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/users [post]
func (u *userHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.UserCreate

		userCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		user := &models.User{
			Firstname: userCreate.Firstname,
			Lastname:  userCreate.Lastname,
			Email:     userCreate.Email,
			Password:  userCreate.Password,
			UserKind:  userCreate.UserKind,
			SchoolId:  userCreate.SchoolId,
		}
		err = user.HashPassword()
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		userDb, err := u.userUseCase.Create(user)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, userDb)
	}
}

// Send Reset Mail
//
//	@Summary		Get all users
//	@Description	Get all users
//	@Tags			User
//	@Produce		json
//	@Success		204
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/users/reset-link [post]
func (u *userHandlers) SendResetMail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.SendResetMail

		userReset, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		resetCode, err := u.userUseCase.SendResetMail(userReset.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		emailM := email.InitEmailManager(u.cfg.Smtp.Username, u.cfg.Smtp.Password, u.cfg.Smtp.Host)
		err = emailM.SendResetEmail([]string{userReset.Email}, resetCode)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}
