package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/document/mock"
	"github.com/esgi-challenge/backend/internal/models"
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
	useCase := NewDocumentUseCase(nil, mockRepo, logger)

	document := &models.Document{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().Create(document).Return(document, nil)

	createdDocument, err := useCase.Create(document)

	assert.NoError(t, err)
	assert.NotNil(t, createdDocument)
}

func TestGetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewDocumentUseCase(nil, mockRepo, logger)

	document := &models.Document{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().GetById(document.ID).Return(document, nil)

	dbDocument, err := useCase.GetById(document.ID)

	assert.NoError(t, err)
	assert.NotNil(t, dbDocument)
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewDocumentUseCase(nil, mockRepo, logger)

	documents := &[]models.Document{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	mockRepo.EXPECT().GetAll().Return(documents, nil)

	dbDocuments, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, dbDocuments)
	assert.Len(t, *dbDocuments, 2)
	assert.Equal(t, &(*documents)[0], &(*dbDocuments)[0])
	assert.Equal(t, &(*documents)[1], &(*dbDocuments)[1])
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewDocumentUseCase(nil, mockRepo, logger)

	document := &models.Document{
		Title:       "title",
		Description: "description",
	}

	t.Run("Delete", func(t *testing.T) {
		mockRepo.EXPECT().GetById(document.ID).Return(document, nil)
		mockRepo.EXPECT().Delete(document.ID).Return(nil)

		err := useCase.Delete(document.ID)

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
	useCase := NewDocumentUseCase(nil, mockRepo, logger)

	document := &models.Document{
		Title:       "title",
		Description: "description",
	}

	fixedDocument := &models.Document{
		Title:       "title updated",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mockRepo.EXPECT().GetById(document.ID).Return(document, nil)
		mockRepo.EXPECT().Update(document.ID, fixedDocument).Return(fixedDocument, nil)

		updatedDocument, err := useCase.Update(document.ID, fixedDocument)

		assert.NoError(t, err)
		assert.NotNil(t, updatedDocument)
		assert.Equal(t, updatedDocument, fixedDocument)
	})

	t.Run("Update Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		updatedDocument, err := useCase.Update(uint(10), fixedDocument)

		assert.Error(t, err)
		assert.Nil(t, updatedDocument)
		assert.EqualError(t, err, "Not found")
	})
}