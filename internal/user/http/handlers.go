package http

import (
	"net/http"

	"github.com/esgi-challenge/backend/config"
	_ "github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
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
