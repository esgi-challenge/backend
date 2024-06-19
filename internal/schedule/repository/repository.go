package repository

import (
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/schedule"
	"gorm.io/gorm"
)

type scheduleRepo struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) schedule.Repository {
	return &scheduleRepo{db: db}
}

func (r *scheduleRepo) Create(schedule *models.Schedule) (*models.Schedule, error) {
	if err := r.db.Create(schedule).Error; err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *scheduleRepo) GetAll() (*[]models.Schedule, error) {
	var schedules []models.Schedule

	if err := r.db.Find(&schedules).Error; err != nil {
		return nil, err
	}

	return &schedules, nil
}

func (r *scheduleRepo) GetById(id uint) (*models.Schedule, error) {
	var schedule models.Schedule

	if err := r.db.First(&schedule, id).Error; err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (r *scheduleRepo) Update(id uint, schedule *models.Schedule) (*models.Schedule, error) {
	if err := r.db.Save(schedule).Error; err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *scheduleRepo) Delete(id uint) error {
	if err := r.db.Debug().Delete(&models.Schedule{}, id).Error; err != nil {
		return err
	}

	return nil
}
