package repository

import (
	"github.com/esgi-challenge/backend/internal/informations"
	"github.com/esgi-challenge/backend/internal/models"
	"gorm.io/gorm"
)

type informationsRepo struct {
	db *gorm.DB
}

func NewInformationsRepository(db *gorm.DB) informations.Repository {
	return &informationsRepo{db: db}
}

func (r *informationsRepo) Create(informations *models.Informations) (*models.Informations, error) {
	if err := r.db.Create(informations).Error; err != nil {
		return nil, err
	}

	return informations, nil
}

func (r *informationsRepo) GetAll() (*[]models.Informations, error) {
	var informationss []models.Informations

	if err := r.db.Find(&informationss).Error; err != nil {
		return nil, err
	}

	return &informationss, nil
}

func (r *informationsRepo) GetBySchoolId(schoolId uint) (*[]models.Informations, error) {
	var informations []models.Informations

	if err := r.db.Find(&informations, "school_id = ?", schoolId).Error; err != nil {
		return nil, err
	}

	return &informations, nil
}

func (r *informationsRepo) GetBySchoolIdAndId(schoolId uint, informationId uint) (*models.Informations, error) {
	var informations models.Informations

	if err := r.db.Find(&informations, "school_id = ? AND id = ", schoolId, informationId).Error; err != nil {
		return nil, err
	}

	return &informations, nil
}

func (r *informationsRepo) GetById(id uint) (*models.Informations, error) {
	var informations models.Informations

	if err := r.db.First(&informations, id).Error; err != nil {
		return nil, err
	}

	return &informations, nil
}

func (r *informationsRepo) Delete(id uint) error {
	if err := r.db.Debug().Delete(&models.Informations{}, id).Error; err != nil {
		return err
	}

	return nil
}
