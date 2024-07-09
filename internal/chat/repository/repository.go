package repository

import (
	"github.com/esgi-challenge/backend/internal/chat"
	"github.com/esgi-challenge/backend/internal/models"
	"gorm.io/gorm"
)

type chatRepo struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) chat.Repository {
	return &chatRepo{db: db}
}

func (r *chatRepo) Create(channel *models.Channel) (*models.Channel, error) {
	if err := r.db.Create(channel).Error; err != nil {
		return nil, err
	}

	return channel, nil
}

func (r *chatRepo) GetAllByUser(userId uint) (*[]models.Channel, error) {
	var channels []models.Channel

	if err := r.db.Model(&models.Channel{}).Preload("FirstUser").Preload("SecondUser").Preload("Messages").Where("first_user_id = ? OR second_user_id = ?", userId, userId).Find(&channels).Error; err != nil {
		return nil, err
	}

	return &channels, nil
}

func (r *chatRepo) GetById(id uint) (*models.Channel, error) {
	var channel models.Channel

	if err := r.db.Model(&models.Channel{}).Preload("FirstUser").Preload("SecondUser").Preload("Messages").First(&channel, id).Error; err != nil {
		return nil, err
	}

	return &channel, nil
}
