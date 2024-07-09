package chat

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(chat *models.Channel) (*models.Channel, error)
	GetAllByUser(userId uint) (*[]models.Channel, error)
	GetById(id uint) (*models.Channel, error)
}
