package repository

import (
	"fmt"

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

func (r *scheduleRepo) Sign(scheduleSignature *models.ScheduleSignature) (*models.ScheduleSignature, error) {
	if err := r.db.Create(scheduleSignature).Error; err != nil {
		return nil, err
	}

	return scheduleSignature, nil
}

func (r *scheduleRepo) GetSign(userId uint, scheduleId uint) (*models.ScheduleSignature, error) {
	var signature models.ScheduleSignature

	if err := r.db.Where("student_id = ?", userId).Where("schedule_id = ?", scheduleId).First(&signature).Error; err != nil {
		return nil, err
	}

	return &signature, nil
}

func (r *scheduleRepo) GetAllBySchoolId(schoolId uint) (*[]models.Schedule, error) {
	var schedules []models.Schedule

	if err := r.db.Model(&models.Schedule{}).Preload("Course").Preload("Campus").Preload("Class").Where("school_id = ?", schoolId).Find(&schedules).Error; err != nil {
		return nil, err
	}

	return &schedules, nil
}

func (r *scheduleRepo) GetAll(userId uint) (*[]models.Schedule, error) {
	var schedules []models.Schedule

	if err := r.db.Raw(getAllByUser, userId).Scan(&schedules).Error; err != nil {
		return nil, err
	}

	return &schedules, nil
}

func (r *scheduleRepo) GetPreloadById(scheduleId uint) (*models.Schedule, error) {
	var schedule models.Schedule

	if err := r.db.Model(&models.Schedule{}).Preload("Course").Preload("Campus").Preload("Class").First(&schedule, scheduleId).Error; err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (r *scheduleRepo) GetById(classRefer, id uint) (*models.Schedule, error) {
	var schedule models.Schedule

	if err := r.db.Model(&models.Schedule{}).Preload("Course").Preload("Course.Teacher").Preload("Campus").Preload("Class").Joins("left join classes on classes.id = class").Where("classes.id = ?", classRefer).First(&schedule, id).Error; err != nil {
		return nil, err
	}

	fmt.Println(schedule)

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
