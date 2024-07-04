package school

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(school *models.School) (*models.School, error)
	GetByUser(user *models.User) (*models.School, error)
	GetAll() (*[]models.School, error)
	GetById(id uint) (*models.School, error)
	Delete(id uint) error
	GetSchoolStudents(schoolId uint) (*[]models.User, error)
	GetSchoolTeachers(schoolId uint) (*[]models.User, error)
}
