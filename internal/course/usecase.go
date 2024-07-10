package course

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, course *models.Course) (*models.Course, error)
	GetAll() (*[]models.Course, error)
	GetAllBySchoolId(schoolId uint) (*[]models.Course, error)
	GetById(id uint) (*models.Course, error)
	Update(id uint, updatedCourse *models.Course) (*models.Course, error)
	Delete(id uint) error
}
