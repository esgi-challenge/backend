package repository

import (
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/user"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) (*models.User, error) {
	return user, nil
}
