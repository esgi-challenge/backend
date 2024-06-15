package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/campus/mock"
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
	useCase := NewCampusUseCase(nil, mockRepo, logger)

	campus := &models.Campus{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().Create(campus).Return(campus, nil)

	createdCampus, err := useCase.Create(campus)

	assert.NoError(t, err)
	assert.NotNil(t, createdCampus)
}

func TestGetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewCampusUseCase(nil, mockRepo, logger)

	campus := &models.Campus{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().GetById(campus.ID).Return(campus, nil)

	dbCampus, err := useCase.GetById(campus.ID)

	assert.NoError(t, err)
	assert.NotNil(t, dbCampus)
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewCampusUseCase(nil, mockRepo, logger)

	campuss := &[]models.Campus{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	mockRepo.EXPECT().GetAll().Return(campuss, nil)

	dbCampuss, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, dbCampuss)
	assert.Len(t, *dbCampuss, 2)
	assert.Equal(t, &(*campuss)[0], &(*dbCampuss)[0])
	assert.Equal(t, &(*campuss)[1], &(*dbCampuss)[1])
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewCampusUseCase(nil, mockRepo, logger)

	campus := &models.Campus{
		Title:       "title",
		Description: "description",
	}

	t.Run("Delete", func(t *testing.T) {
		mockRepo.EXPECT().GetById(campus.ID).Return(campus, nil)
		mockRepo.EXPECT().Delete(campus.ID).Return(nil)

		err := useCase.Delete(campus.ID)

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
	useCase := NewCampusUseCase(nil, mockRepo, logger)

	campus := &models.Campus{
		Title:       "title",
		Description: "description",
	}

	fixedCampus := &models.Campus{
		Title:       "title updated",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mockRepo.EXPECT().GetById(campus.ID).Return(campus, nil)
		mockRepo.EXPECT().Update(campus.ID, fixedCampus).Return(fixedCampus, nil)

		updatedCampus, err := useCase.Update(campus.ID, fixedCampus)

		assert.NoError(t, err)
		assert.NotNil(t, updatedCampus)
		assert.Equal(t, updatedCampus, fixedCampus)
	})

	t.Run("Update Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		updatedCampus, err := useCase.Update(uint(10), fixedCampus)

		assert.Error(t, err)
		assert.Nil(t, updatedCampus)
		assert.EqualError(t, err, "Not found")
	})
}
