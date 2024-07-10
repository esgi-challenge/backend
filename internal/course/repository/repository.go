package repository

import (
	"github.com/esgi-challenge/backend/internal/course"
	"github.com/esgi-challenge/backend/internal/models"
	"gorm.io/gorm"
)

type courseRepo struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) course.Repository {
	return &courseRepo{db: db}
}

func (r *courseRepo) Create(course *models.Course) (*models.Course, error) {
	if err := r.db.Create(course).Error; err != nil {
		return nil, err
	}

	return course, nil
}

func (r *courseRepo) GetAll() (*[]models.Course, error) {
	var courses []models.Course

	if err := r.db.Find(&courses).Error; err != nil {
		return nil, err
	}

	return &courses, nil
}

func (r *courseRepo) GetAllBySchoolId(schoolId uint) (*[]models.Course, error) {
	var courses []models.Course

	if err := r.db.Model(&models.Course{}).Preload("Teacher").Preload("Path").Where("school_id = ?", schoolId).Find(&courses).Error; err != nil {
		return nil, err
	}

	return &courses, nil
}

func (r *courseRepo) GetById(id uint) (*models.Course, error) {
	var course models.Course

	if err := r.db.Model(&models.Course{}).Preload("Teacher").Preload("Path").First(&course, id).Error; err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *courseRepo) Update(id uint, course *models.Course) (*models.Course, error) {
	if err := r.db.Save(course).Error; err != nil {
		return nil, err
	}

	return course, nil
}

func (r *courseRepo) Delete(id uint) error {
	if err := r.db.Debug().Delete(&models.Course{}, id).Error; err != nil {
		return err
	}

	return nil
}
