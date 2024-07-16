package note

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(note *models.Note) (*models.Note, error)
	GetAllByStudent(studentId uint) (*[]models.Note, error)
	GetAllByTeacher(teacherId uint) (*[]models.Note, error)
	GetById(id uint) (*models.Note, error)
	GetByIdPreload(id uint) (*models.Note, error)
	Update(id uint, note *models.Note) (*models.Note, error)
	Delete(id uint) error
}
