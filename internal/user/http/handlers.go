package http

import (
	"fmt"
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/user"
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

func (u *userHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := &models.User{
			Username: "admin",
			Email:    "admin@admin.fr",
			Password: "password",
		}
		createdUser, err := u.userUseCase.Create(ctx, user)
		if err != nil {
			fmt.Print("Error")
		}
		fmt.Print(createdUser)

		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}
