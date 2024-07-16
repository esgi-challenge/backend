package repository

import (
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/note"
	"gorm.io/gorm"
)

type noteRepo struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) note.Repository {
	return &noteRepo{db: db}
}

func (r *noteRepo) Create(note *models.Note) (*models.Note, error) {
	if err := r.db.Create(note).Error; err != nil {
		return nil, err
	}

	return note, nil
}

func (r *noteRepo) GetAllByStudent(studentId uint) (*[]models.Note, error) {
	var notes []models.Note

	if err := r.db.Model(&models.Note{}).Preload("Project.Course").Preload("Teacher").Find(&notes).Where("student_id = ?", studentId).Error; err != nil {
		return nil, err
	}

	return &notes, nil
}

func (r *noteRepo) GetAllByTeacher(teacherId uint) (*[]models.Note, error) {
	var notes []models.Note

	if err := r.db.Model(&models.Note{}).Preload("Project").Preload("Student").Find(&notes).Where("teacher_id = ?", teacherId).Error; err != nil {
		return nil, err
	}

	return &notes, nil
}

func (r *noteRepo) GetById(id uint) (*models.Note, error) {
	var note models.Note

	if err := r.db.First(&note, id).Error; err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *noteRepo) GetByIdPreload(id uint) (*models.Note, error) {
	var note models.Note

	if err := r.db.Model(&models.Note{}).Preload("Project").Preload("Student").First(&note, id).Error; err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *noteRepo) Update(id uint, note *models.Note) (*models.Note, error) {
	if err := r.db.Save(note).Error; err != nil {
		return nil, err
	}

	return note, nil
}

func (r *noteRepo) Delete(id uint) error {
	if err := r.db.Debug().Delete(&models.Note{}, id).Error; err != nil {
		return err
	}

	return nil
}
