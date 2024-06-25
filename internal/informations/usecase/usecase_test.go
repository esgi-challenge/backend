package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/informations/mock"
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
	useCase := NewInformationsUseCase(nil, mockRepo, logger)

	informations := &models.Informations{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().Create(informations).Return(informations, nil)

	createdInformations, err := useCase.Create(informations)

	assert.NoError(t, err)
	assert.NotNil(t, createdInformations)
}

func TestGetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewInformationsUseCase(nil, mockRepo, logger)

	informations := &models.Informations{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().GetById(informations.ID).Return(informations, nil)

	dbInformations, err := useCase.GetById(informations.ID)

	assert.NoError(t, err)
	assert.NotNil(t, dbInformations)
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewInformationsUseCase(nil, mockRepo, logger)

	informationss := &[]models.Informations{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	mockRepo.EXPECT().GetAll().Return(informationss, nil)

	dbInformationss, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, dbInformationss)
	assert.Len(t, *dbInformationss, 2)
	assert.Equal(t, &(*informationss)[0], &(*dbInformationss)[0])
	assert.Equal(t, &(*informationss)[1], &(*dbInformationss)[1])
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewInformationsUseCase(nil, mockRepo, logger)

	informations := &models.Informations{
		Title:       "title",
		Description: "description",
	}

	t.Run("Delete", func(t *testing.T) {
		mockRepo.EXPECT().GetById(informations.ID).Return(informations, nil)
		mockRepo.EXPECT().Delete(informations.ID).Return(nil)

		err := useCase.Delete(informations.ID)

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
	useCase := NewInformationsUseCase(nil, mockRepo, logger)

	informations := &models.Informations{
		Title:       "title",
		Description: "description",
	}

	fixedInformations := &models.Informations{
		Title:       "title updated",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mockRepo.EXPECT().GetById(informations.ID).Return(informations, nil)
		mockRepo.EXPECT().Update(informations.ID, fixedInformations).Return(fixedInformations, nil)

		updatedInformations, err := useCase.Update(informations.ID, fixedInformations)

		assert.NoError(t, err)
		assert.NotNil(t, updatedInformations)
		assert.Equal(t, updatedInformations, fixedInformations)
	})

	t.Run("Update Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		updatedInformations, err := useCase.Update(uint(10), fixedInformations)

		assert.Error(t, err)
		assert.Nil(t, updatedInformations)
		assert.EqualError(t, err, "Not found")
	})
}