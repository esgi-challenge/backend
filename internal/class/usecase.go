package class

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, class *models.Class) (*models.Class, error)
	Add(user *models.User, id uint, class *models.ClassAdd) (*models.Class, error)
	GetAll() (*[]models.Class, error)
	GetById(id uint) (*models.Class, error)
	Update(user *models.User, id uint, updatedClass *models.Class) (*models.Class, error)
	Delete(user *models.User, id uint) error
}
