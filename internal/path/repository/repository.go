package repository

import (
	"github.com/esgi-challenge/backend/internal/path"
	"github.com/esgi-challenge/backend/internal/models"
	"gorm.io/gorm"
)

type pathRepo struct {
	db *gorm.DB
}

func NewPathRepository(db *gorm.DB) path.Repository {
	return &pathRepo{db: db}
}

func (r *pathRepo) Create(path *models.Path) (*models.Path, error) {
	if err := r.db.Create(path).Error; err != nil {
		return nil, err
	}

	return path, nil
}

func (r *pathRepo) GetAll() (*[]models.Path, error) {
	var paths []models.Path

	if err := r.db.Find(&paths).Error; err != nil {
		return nil, err
	}

	return &paths, nil
}

func (r *pathRepo) GetById(id uint) (*models.Path, error) {
	var path models.Path

	if err := r.db.First(&path, id).Error; err != nil {
		return nil, err
	}

	return &path, nil
}

func (r *pathRepo) Update(id uint, path *models.Path) (*models.Path, error) {
	if err := r.db.Save(path).Error; err != nil {
		return nil, err
	}

	return path, nil
}

func (r *pathRepo) Delete(id uint) error {
	if err := r.db.Debug().Delete(&models.Path{}, id).Error; err != nil {
		return err
	}

	return nil
}