package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/note/mock"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewNoteUseCase(nil, mockRepo, logger)

	note := &models.Note{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().Create(note).Return(note, nil)

	createdNote, err := useCase.Create(note)

	assert.NoError(t, err)
	assert.NotNil(t, createdNote)
}

func TestGetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewNoteUseCase(nil, mockRepo, logger)

	note := &models.Note{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().GetById(note.ID).Return(note, nil)

	dbNote, err := useCase.GetById(note.ID)

	assert.NoError(t, err)
	assert.NotNil(t, dbNote)
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewNoteUseCase(nil, mockRepo, logger)

	notes := &[]models.Note{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	mockRepo.EXPECT().GetAll().Return(notes, nil)

	dbNotes, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, dbNotes)
	assert.Len(t, *dbNotes, 2)
	assert.Equal(t, &(*notes)[0], &(*dbNotes)[0])
	assert.Equal(t, &(*notes)[1], &(*dbNotes)[1])
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewNoteUseCase(nil, mockRepo, logger)

	note := &models.Note{
		Title:       "title",
		Description: "description",
	}

	t.Run("Delete", func(t *testing.T) {
		mockRepo.EXPECT().GetById(note.ID).Return(note, nil)
		mockRepo.EXPECT().Delete(note.ID).Return(nil)

		err := useCase.Delete(note.ID)

		assert.NoError(t, err)
	})

	t.Run("Delete Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		err := useCase.Delete(uint(10))

		assert.Error(t, err)
		assert.EqualError(t, err, "Not found")
	})
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewNoteUseCase(nil, mockRepo, logger)

	note := &models.Note{
		Title:       "title",
		Description: "description",
	}

	fixedNote := &models.Note{
		Title:       "title updated",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mockRepo.EXPECT().GetById(note.ID).Return(note, nil)
		mockRepo.EXPECT().Update(note.ID, fixedNote).Return(fixedNote, nil)

		updatedNote, err := useCase.Update(note.ID, fixedNote)

		assert.NoError(t, err)
		assert.NotNil(t, updatedNote)
		assert.Equal(t, updatedNote, fixedNote)
	})

	t.Run("Update Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		updatedNote, err := useCase.Update(uint(10), fixedNote)

		assert.Error(t, err)
		assert.Nil(t, updatedNote)
		assert.EqualError(t, err, "Not found")
	})
}
