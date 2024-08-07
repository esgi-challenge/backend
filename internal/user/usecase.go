package user

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User) (*models.User, error)
	GetAll() (*[]models.User, error)
	SendResetMail(email string) (string, error)
	Update(id uint, updatedUser *models.User) (*models.User, error)
	GetById(id uint) (*models.User, error)
}
