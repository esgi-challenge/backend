package class

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(class *models.Class) (*models.Class, error)
	GetAll() (*[]models.Class, error)
	GetById(id uint) (*models.Class, error)
	Update(id uint, class *models.Class) (*models.Class, error)
	Delete(id uint) error
}
