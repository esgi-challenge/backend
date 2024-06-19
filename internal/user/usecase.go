package user

import (
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/gin-gonic/gin"
)

type UseCase interface {
	Create(ctx *gin.Context, user *models.User) (*models.User, error)
	GetAll() (*[]models.User, error)
}
