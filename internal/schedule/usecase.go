package schedule

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, schedule *models.ScheduleCreate) (*models.Schedule, error)
	GetAll(user *models.User) (*[]models.ScheduleGet, error)
	GetUnattended(user *models.User) ([]models.ScheduleGet, error)
	GetAllBySchoolId(schoolId uint) (*[]models.Schedule, error)
	GetById(user *models.User, id uint) (*models.ScheduleGet, error)
	GetPreloadById(scheduleId uint) (*models.Schedule, error)
	Sign(signature *models.ScheduleSignatureCreate, user *models.User, id uint) (*models.ScheduleSignature, error)
	CheckSign(user *models.User, id uint) (*models.ScheduleSignature, error)
	GetSignatureCode(user *models.User, scheduleId uint) (*models.ScheduleSignatureCode, error)
	Update(user *models.User, id uint, updatedSchedule *models.Schedule) (*models.Schedule, error)
	Delete(user *models.User, id uint) error
	GetStudentsSignature(user *models.User, scheduleId uint) (*models.ScheduleSignatureGet, error)
}
