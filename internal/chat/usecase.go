package chat

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(chat *models.Channel) (*models.Channel, error)
	GetAllByUser(userId uint) (*[]models.Channel, error)
	GetById(id uint) (*models.Channel, error)
	SaveMessage(msg *models.Message) (*models.Message, error)
	GetAllPossibleChatStudent(user *models.User) (*[]models.User, error)
	GetAllPossibleChatTeacher(user *models.User) (*[]models.User, error)
}
