package school

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, school *models.SchoolCreate) (*models.School, error)
	Invite(user *models.User, school *models.SchoolInvite) (*models.User, error)

	GetAll() (*[]models.School, error)
	GetById(id uint) (*models.School, error)
  GetByUser(user *models.User) (*models.School, error)
	Delete(user *models.User, id uint) error
  GetSchoolStudents(adminID uint) (*[]models.User, error)
}
