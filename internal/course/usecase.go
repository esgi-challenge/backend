package course

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, course *models.Course) (*models.Course, error)
	GetAll() (*[]models.Course, error)
	GetById(id uint) (*models.Course, error)
	Update(user *models.User, id uint, updatedCourse *models.Course) (*models.Course, error)
	Delete(user *models.User, id uint) error
}
