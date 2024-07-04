package path

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, path *models.Path) (*models.Path, error)
	GetAll() (*[]models.Path, error)
	GetAllBySchoolId(schoolId uint) (*[]models.Path, error)
	GetById(id uint) (*models.Path, error)
	Update(id uint, updatedPath *models.Path) (*models.Path, error)
	Delete(id uint) error
}
