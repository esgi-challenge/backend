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
  GetSchoolStudentsByAdminID(adminID uint) (*[]models.User, error)
}
