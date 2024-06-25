package informations

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, informations *models.Informations) (*models.Informations, error)
	GetAll(user *models.User) (*[]models.Informations, error)
	GetById(user *models.User, id uint) (*models.Informations, error)
	Update(user *models.User, id uint, updatedInformations *models.Informations) (*models.Informations, error)
	Delete(user *models.User, id uint) error
}
