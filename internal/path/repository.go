package path

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(path *models.Path) (*models.Path, error)
	GetAll() (*[]models.Path, error)
	GetById(id uint) (*models.Path, error)
	Update(id uint, path *models.Path) (*models.Path, error)
	Delete(id uint) error
}
