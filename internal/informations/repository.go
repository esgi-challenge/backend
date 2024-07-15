package informations

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(informations *models.Informations) (*models.Informations, error)
	GetAll() (*[]models.Informations, error)
	GetById(id uint) (*models.Informations, error)
	GetBySchoolId(schoolId uint) (*[]models.Informations, error)
	GetBySchoolIdAndId(schoolId uint, informationId uint) (*models.Informations, error)
	Delete(id uint) error
}
