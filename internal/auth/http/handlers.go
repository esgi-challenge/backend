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

// Create
//
//	@Summary		Log to the api
//	@Description	Log to the api
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			auth	body		models.AuthLogin	true	"Login Credentials"
//	@Success		201		{object}	models.Auth
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/auth [post]
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

		ctx.JSON(http.StatusCreated, token)
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
//	@Router			/auth [post]
func (u *authHandlers) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.AuthRegister

		authRegister, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		payload := &models.AuthRegister{
			Firstname: authRegister.Firstname,
			Lastname:  authRegister.Lastname,
			Email:     authRegister.Email,
			Password:  authRegister.Password,
		}

		token, err := u.authUseCase.Register(payload)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, token)
	}
}
