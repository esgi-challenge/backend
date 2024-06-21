package user

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(user *models.User) (*models.User, error)
	GetAll() (*[]models.User, error)
	GetById(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByInvitationCode(invitationCode string) (*models.User, error)
	GetByResetCode(resetCode string) (*models.User, error)
	Update(id uint, user *models.User) (*models.User, error)
}
