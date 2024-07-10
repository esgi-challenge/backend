package document

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(document *models.Document) (*models.Document, error)
	GetById(id uint) (*models.Document, error)
	Delete(id uint) error
}
