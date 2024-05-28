package school

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(school *models.School) (*models.School, error)
	GetAll() (*[]models.School, error)
	GetById(id uint) (*models.School, error)
	Delete(id uint) error
}
