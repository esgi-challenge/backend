package repository

import (
	"github.com/esgi-challenge/backend/internal/campus"
	"github.com/esgi-challenge/backend/internal/models"
	"gorm.io/gorm"
)

type campusRepo struct {
	db *gorm.DB
}

func NewCampusRepository(db *gorm.DB) campus.Repository {
	return &campusRepo{db: db}
}

func (r *campusRepo) Create(campus *models.Campus) (*models.Campus, error) {
	if err := r.db.Create(campus).Error; err != nil {
		return nil, err
	}

	return campus, nil
}

func (r *campusRepo) GetAll() (*[]models.Campus, error) {
	var campuss []models.Campus

	if err := r.db.Find(&campuss).Error; err != nil {
		return nil, err
	}

	return &campuss, nil
}

func (r *campusRepo) GetAllBySchoolId(schoolId uint) (*[]models.Campus, error) {
	var campus []models.Campus

	if err := r.db.Where("school_id = ?", schoolId).Find(&campus).Error; err != nil {
		return nil, err
	}

	return &campus, nil
}

func (r *campusRepo) GetById(id uint) (*models.Campus, error) {
	var campus models.Campus

	if err := r.db.First(&campus, id).Error; err != nil {
		return nil, err
	}

	return &campus, nil
}

func (r *campusRepo) Update(id uint, campus *models.Campus) (*models.Campus, error) {
	if err := r.db.Save(campus).Error; err != nil {
		return nil, err
	}

	return campus, nil
}

func (r *campusRepo) Delete(id uint) error {
	if err := r.db.Debug().Delete(&models.Campus{}, id).Error; err != nil {
		return err
	}

	return nil
}
