package course

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(course *models.Course) (*models.Course, error)
	GetAll() (*[]models.Course, error)
	GetById(id uint) (*models.Course, error)
	Update(id uint, course *models.Course) (*models.Course, error)
	Delete(id uint) error
	GetAllBySchoolId(schoolId uint) (*[]models.Course, error)
}
