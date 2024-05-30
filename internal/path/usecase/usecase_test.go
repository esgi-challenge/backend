package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/path/mock"
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
	useCase := NewPathUseCase(nil, mockRepo, logger)

	path := &models.Path{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().Create(path).Return(path, nil)

	createdPath, err := useCase.Create(path)

	assert.NoError(t, err)
	assert.NotNil(t, createdPath)
}

func TestGetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewPathUseCase(nil, mockRepo, logger)

	path := &models.Path{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().GetById(path.ID).Return(path, nil)

	dbPath, err := useCase.GetById(path.ID)

	assert.NoError(t, err)
	assert.NotNil(t, dbPath)
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewPathUseCase(nil, mockRepo, logger)

	paths := &[]models.Path{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	mockRepo.EXPECT().GetAll().Return(paths, nil)

	dbPaths, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, dbPaths)
	assert.Len(t, *dbPaths, 2)
	assert.Equal(t, &(*paths)[0], &(*dbPaths)[0])
	assert.Equal(t, &(*paths)[1], &(*dbPaths)[1])
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewPathUseCase(nil, mockRepo, logger)

	path := &models.Path{
		Title:       "title",
		Description: "description",
	}

	t.Run("Delete", func(t *testing.T) {
		mockRepo.EXPECT().GetById(path.ID).Return(path, nil)
		mockRepo.EXPECT().Delete(path.ID).Return(nil)

		err := useCase.Delete(path.ID)

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
	useCase := NewPathUseCase(nil, mockRepo, logger)

	path := &models.Path{
		Title:       "title",
		Description: "description",
	}

	fixedPath := &models.Path{
		Title:       "title updated",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mockRepo.EXPECT().GetById(path.ID).Return(path, nil)
		mockRepo.EXPECT().Update(path.ID, fixedPath).Return(fixedPath, nil)

		updatedPath, err := useCase.Update(path.ID, fixedPath)

		assert.NoError(t, err)
		assert.NotNil(t, updatedPath)
		assert.Equal(t, updatedPath, fixedPath)
	})

	t.Run("Update Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		updatedPath, err := useCase.Update(uint(10), fixedPath)

		assert.Error(t, err)
		assert.Nil(t, updatedPath)
		assert.EqualError(t, err, "Not found")
	})
}