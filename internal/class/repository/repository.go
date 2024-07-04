package repository

import (
	"github.com/esgi-challenge/backend/internal/class"
	"github.com/esgi-challenge/backend/internal/models"
	"gorm.io/gorm"
)

type classRepo struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) class.Repository {
	return &classRepo{db: db}
}

func (r *classRepo) Create(class *models.Class) (*models.Class, error) {
	if err := r.db.Create(class).Error; err != nil {
		return nil, err
	}

	return class, nil
}

func (r *classRepo) GetAll() (*[]models.Class, error) {
	var classs []models.Class

	if err := r.db.Model(&models.Class{}).Preload("Students").Find(&classs).Error; err != nil {
		return nil, err
	}

	return &classs, nil
}

func (r *classRepo) GetAllBySchoolId(schoolId uint) (*[]models.Class, error) {
	var classes []models.Class

	if err := r.db.Model(&models.Class{}).Preload("Students").Where("school_id = ?", schoolId).Find(&classes).Error; err != nil {
		return nil, err
	}

	return &classes, nil
}

func (r *classRepo) GetById(id uint) (*models.Class, error) {
	var class models.Class

	if err := r.db.Model(&models.Class{}).Preload("Students").First(&class, id).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (r *classRepo) Update(id uint, class *models.Class) (*models.Class, error) {
	if err := r.db.Save(class).Error; err != nil {
		return nil, err
	}

	return class, nil
}

func (r *classRepo) Delete(id uint) error {
	if err := r.db.Debug().Delete(&models.Class{}, id).Error; err != nil {
		return err
	}

	return nil
}
