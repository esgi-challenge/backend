package schedule

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(schedule *models.Schedule) (*models.Schedule, error)
	Sign(schedule *models.ScheduleSignature) (*models.ScheduleSignature, error)
	GetSign(userId uint, scheduleId uint) (*models.ScheduleSignature, error)
	GetAllBySchoolId(schoolId uint) (*[]models.Schedule, error)
	GetPreloadById(scheduleId uint) (*models.Schedule, error)
	GetById(userId, id uint) (*models.Schedule, error)
	Update(id uint, schedule *models.Schedule) (*models.Schedule, error)
	Delete(id uint) error
}
