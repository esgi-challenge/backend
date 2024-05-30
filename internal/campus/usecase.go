package campus

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, campus *models.Campus) (*models.Campus, error)
	GetAll() (*[]models.Campus, error)
	GetById(id uint) (*models.Campus, error)
	Update(user *models.User, id uint, updatedCampus *models.Campus) (*models.Campus, error)
	Delete(user *models.User, id uint) error
}
