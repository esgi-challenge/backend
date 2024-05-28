package user

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(user *models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}
