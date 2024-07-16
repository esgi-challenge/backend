package usecase

import (
	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/note"
	"github.com/esgi-challenge/backend/pkg/logger"
	"gorm.io/gorm"
)

type noteUseCase struct {
	noteRepo note.Repository
	cfg      *config.Config
	logger   logger.Logger
}

func NewNoteUseCase(cfg *config.Config, noteRepo note.Repository, logger logger.Logger) note.UseCase {
	return &noteUseCase{cfg: cfg, noteRepo: noteRepo, logger: logger}
}

func (u *noteUseCase) Create(note *models.Note) (*models.Note, error) {
	note, err := u.noteRepo.Create(note)
	if err != nil {
		return nil, err
	}

	return u.noteRepo.GetByIdPreload(note.ID)
}

func (u *noteUseCase) GetAllByUser(user *models.User) (*[]models.Note, error) {
	if *user.UserKind == models.TEACHER {
		return u.noteRepo.GetAllByTeacher(user.ID)
	} else if *user.UserKind == models.STUDENT {
		return u.noteRepo.GetAllByStudent(user.ID)
	} else {
		return nil, gorm.ErrRecordNotFound
	}
}

func (u *noteUseCase) GetById(id uint) (*models.Note, error) {
	return u.noteRepo.GetById(id)
}

func (u *noteUseCase) Update(id uint, updatedNote *models.Note) (*models.Note, error) {
	// Temporary fix for known issue :
	// https://github.com/go-gorm/gorm/issues/5724
	//////////////////////////////////////
	dbNote, err := u.GetById(id)
	if err != nil {
		return nil, err
	}
	updatedNote.CreatedAt = dbNote.CreatedAt
	///////////////////////////////////////

	updatedNote.ID = id
	note, err := u.noteRepo.Update(id, updatedNote)
	if err != nil {
		return nil, err
	}

	return u.noteRepo.GetByIdPreload(note.ID)
}

func (u *noteUseCase) Delete(id uint) error {
	// Check not needed but added to handle a not found error because gorm do not return
	// error if delete on a row that does not exist
	_, err := u.GetById(id)
	if err != nil {
		return err
	}

	return u.noteRepo.Delete(id)
}
