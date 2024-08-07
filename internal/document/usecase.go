package document

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, document *models.DocumentCreate) (*models.Document, error)
	GetById(user *models.User, id uint) (*models.Document, error)
	GetAllByUserId(userId uint) (*[]models.Document, error)
	GetAll(user *models.User) (*[]models.Document, error)
	Delete(user *models.User, id uint) error
}
