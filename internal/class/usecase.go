package class

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, class *models.Class) (*models.Class, error)
	Add(id uint, class *models.ClassAdd) (*models.User, error)
	Remove(id uint, class *models.ClassRemove) (*models.User, error)
	GetAll() (*[]models.Class, error)
	GetAllBySchoolId(schoolId uint) (*[]models.Class, error)
	GetClassLessStudents(schoolId uint) (*[]models.User, error)
	GetById(id uint) (*models.Class, error)
	Update(id uint, updatedClass *models.Class) (*models.Class, error)
	Delete(user *models.User, id uint) error
}
