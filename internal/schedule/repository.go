package schedule

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(schedule *models.Schedule) (*models.Schedule, error)
	GetAll() (*[]models.Schedule, error)
	GetById(id uint) (*models.Schedule, error)
	Update(id uint, schedule *models.Schedule) (*models.Schedule, error)
	Delete(id uint) error
}
