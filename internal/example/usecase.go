package example

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.Example) (*models.Example, error)
	GetAll() (*[]models.Example, error)
	GetById(id uint) (*models.Example, error)
	Update(id uint, updatedExample *models.Example) (*models.Example, error)
	Delete(id uint) error
}
