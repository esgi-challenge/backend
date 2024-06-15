package schedule

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, schedule *models.Schedule) (*models.Schedule, error)
	GetAll(user *models.User) (*[]models.Schedule, error)
	GetById(user *models.User, id uint) (*models.Schedule, error)
	Update(user *models.User, id uint, updatedSchedule *models.Schedule) (*models.Schedule, error)
	Delete(user *models.User, id uint) error
}
