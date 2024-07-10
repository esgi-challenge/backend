package schedule

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, schedule *models.ScheduleCreate) (*models.Schedule, error)
	GetAll(user *models.User) (*[]models.Schedule, error)
	GetById(user *models.User, id uint) (*models.Schedule, error)
	Sign(signature *models.ScheduleSignatureCreate, user *models.User, id uint) (*models.ScheduleSignature, error)
	GetSignatureCode(user *models.User, scheduleId uint) (*models.ScheduleSignatureCode, error)
	Update(user *models.User, id uint, updatedSchedule *models.Schedule) (*models.Schedule, error)
	Delete(user *models.User, id uint) error
}
