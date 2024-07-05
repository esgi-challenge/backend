package class

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(class *models.Class) (*models.Class, error)
	GetAll() (*[]models.Class, error)
	GetAllBySchoolId(schoolId uint) (*[]models.Class, error)
	GetClassLessStudents(schoolId uint) (*[]models.User, error)
	GetById(id uint) (*models.Class, error)
	Update(id uint, class *models.Class) (*models.Class, error)
	Delete(id uint) error
}
