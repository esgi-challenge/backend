package campus

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(campus *models.Campus) (*models.Campus, error)
	GetAll() (*[]models.Campus, error)
	GetById(id uint) (*models.Campus, error)
	Update(id uint, campus *models.Campus) (*models.Campus, error)
	Delete(id uint) error
}