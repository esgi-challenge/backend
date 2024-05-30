package path

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, path *models.Path) (*models.Path, error)
	GetAll() (*[]models.Path, error)
	GetById(id uint) (*models.Path, error)
	Update(user *models.User, id uint, updatedPath *models.Path) (*models.Path, error)
	Delete(user *models.User, id uint) error
}
