package http

import (
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/auth"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type authHandlers struct {
	cfg         *config.Config
	authUseCase auth.UseCase
	logger      logger.Logger
}

func NewAuthHandlers(cfg *config.Config, authUseCase auth.UseCase, logger logger.Logger) auth.Handlers {
	return &authHandlers{cfg: cfg, authUseCase: authUseCase, logger: logger}
}

// Login
//
//	@Summary		Log to the api
//	@Description	Log to the api
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			auth	body		models.AuthLogin	true	"Login Credentials"
//	@Success		200		{object}	models.Auth
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/auth/login [post]
func (u *authHandlers) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.AuthLogin

		authLogin, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		payload := &models.AuthLogin{
			Email:    authLogin.Email,
			Password: authLogin.Password,
		}

		token, err := u.authUseCase.Login(payload)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, token)
	}
}

// Create
//
//	@Summary		Register to the api
//	@Description	Register to the api
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			auth	body		models.AuthRegister	true	"Register Infos"
//	@Success		201		{object}	models.Auth
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/auth/register [post]
func (u *authHandlers) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.AuthRegister

		authRegister, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		user := &models.User{
			Firstname: authRegister.Firstname,
			Lastname:  authRegister.Lastname,
			Email:     authRegister.Email,
			Password:  authRegister.Password,
			UserKind:  models.ADMINISTRATOR,
		}
		err = user.HashPassword()
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		token, err := u.authUseCase.Register(user)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, token)
	}
}

// Invitation Code
//
//	@Summary		Create Password With AuthCode
//	@Description	Create Password With AuthCode
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			auth	body		models.AuthInvitationCode	true	"Register Infos"
//	@Success		200		{object}	models.Auth
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/auth/invitation-code [post]
func (u *authHandlers) InvitationCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.AuthInvitationCode

		authInvitationCode, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		payload := &models.AuthInvitationCode{
			InvitationCode: authInvitationCode.InvitationCode,
			Password:       authInvitationCode.Password,
		}

		token, err := u.authUseCase.InvitationCode(payload)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, token)
	}
}

// Reset Password
//
//	@Summary		Reset User Password
//	@Description	Reset User Password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			auth	body		models.AuthRegister	true	"Register Infos"
//	@Success		200		{object}	models.Auth
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/auth [post]
func (u *authHandlers) ResetPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.AuthResetPassword

		authResetPassword, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		payload := &models.AuthResetPassword{
			Code:     authResetPassword.Code,
			Password: authResetPassword.Password,
		}

		err = u.authUseCase.ResetPassword(payload)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}
