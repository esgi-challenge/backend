package repository

import (
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/school"
	"gorm.io/gorm"
)

type schoolRepo struct {
	db *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) school.Repository {
	return &schoolRepo{db: db}
}

func (r *schoolRepo) Create(school *models.School) (*models.School, error) {
	if err := r.db.Create(school).Error; err != nil {
		return nil, err
	}

	return school, nil
}

func (r *schoolRepo) GetByUser(user *models.User) (*models.School, error) {
	var school models.School

	if err := r.db.First(&school, "userid = ?", user.ID).Error; err != nil {
		return nil, err
	}

	return &school, nil
}

func (r *schoolRepo) GetAll() (*[]models.School, error) {
	var schools []models.School

	if err := r.db.Find(&schools).Error; err != nil {
		return nil, err
	}

	return &schools, nil
}

func (r *schoolRepo) GetById(id uint) (*models.School, error) {
	var school models.School

	if err := r.db.First(&school, id).Error; err != nil {
		return nil, err
	}

	return &school, nil
}

func (r *schoolRepo) Delete(id uint) error {
	if err := r.db.Debug().Delete(&models.School{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *schoolRepo) GetStudentsBySchoolID(schoolID uint) (*[]models.User, error) {
    var students []models.User

    if err := r.db.Where("school_id = ? AND user_kind = ?", schoolID, models.STUDENT).Find(&students).Error; err != nil {
        return nil, err
    }

    return &students, nil
}
