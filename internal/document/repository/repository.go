package repository

import (
	"github.com/esgi-challenge/backend/internal/document"
	"github.com/esgi-challenge/backend/internal/models"
	"gorm.io/gorm"
)

type documentRepo struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) document.Repository {
	return &documentRepo{db: db}
}

func (r *documentRepo) Create(document *models.Document) (*models.Document, error) {
	if err := r.db.Create(document).Error; err != nil {
		return nil, err
	}

	return document, nil
}

func (r *documentRepo) GetAll() (*[]models.Document, error) {
	var documents []models.Document

	if err := r.db.Preload("Course").Find(&documents).Error; err != nil {
		return nil, err
	}

	return &documents, nil
}

func (r *documentRepo) GetById(id uint) (*models.Document, error) {
	var document models.Document

	if err := r.db.Preload("Course").First(&document, id).Error; err != nil {
		return nil, err
	}

	return &document, nil
}

func (r *documentRepo) GetAllByUserId(userId uint) (*[]models.Document, error) {
	var documents []models.Document

	if err := r.db.Preload("Course").Where("user_id = ?", userId).Find(&documents).Error; err != nil {
		return nil, err
	}

	return &documents, nil
}

func (r *documentRepo) Delete(id uint) error {
	if err := r.db.Debug().Delete(&models.Document{}, id).Error; err != nil {
		return err
	}

	return nil
}
