package note

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(note *models.Note) (*models.Note, error)
	GetAllByUser(user *models.User) (*[]models.Note, error)
	GetById(id uint) (*models.Note, error)
	Update(id uint, updatedNote *models.Note) (*models.Note, error)
	Delete(id uint) error
}
