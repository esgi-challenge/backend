package repository

import (
	"github.com/esgi-challenge/backend/internal/example"
	"github.com/esgi-challenge/backend/internal/models"
	"gorm.io/gorm"
)

type exampleRepo struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) example.Repository {
	return &exampleRepo{db: db}
}

func (r *exampleRepo) Create(example *models.Example) (*models.Example, error) {
	if err := r.db.Create(example).Error; err != nil {
		return nil, err
	}

	return example, nil
}

func (r *exampleRepo) GetAll() (*[]models.Example, error) {
	var examples []models.Example

	if err := r.db.Find(&examples).Error; err != nil {
		return nil, err
	}

	return &examples, nil
}

func (r *exampleRepo) GetById(id uint) (*models.Example, error) {
	var example models.Example

	if err := r.db.First(&example, id).Error; err != nil {
		return nil, err
	}

	return &example, nil
}

func (r *exampleRepo) Delete(id uint) error {
	if err := r.db.Delete(&models.Example{}, id).Error; err != nil {
		return err
	}

	return nil
}
