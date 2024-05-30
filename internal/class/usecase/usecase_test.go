package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/class/mock"
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
	useCase := NewClassUseCase(nil, mockRepo, logger)

	class := &models.Class{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().Create(class).Return(class, nil)

	createdClass, err := useCase.Create(class)

	assert.NoError(t, err)
	assert.NotNil(t, createdClass)
}

func TestGetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewClassUseCase(nil, mockRepo, logger)

	class := &models.Class{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().GetById(class.ID).Return(class, nil)

	dbClass, err := useCase.GetById(class.ID)

	assert.NoError(t, err)
	assert.NotNil(t, dbClass)
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewClassUseCase(nil, mockRepo, logger)

	classs := &[]models.Class{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	mockRepo.EXPECT().GetAll().Return(classs, nil)

	dbClasss, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, dbClasss)
	assert.Len(t, *dbClasss, 2)
	assert.Equal(t, &(*classs)[0], &(*dbClasss)[0])
	assert.Equal(t, &(*classs)[1], &(*dbClasss)[1])
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewClassUseCase(nil, mockRepo, logger)

	class := &models.Class{
		Title:       "title",
		Description: "description",
	}

	t.Run("Delete", func(t *testing.T) {
		mockRepo.EXPECT().GetById(class.ID).Return(class, nil)
		mockRepo.EXPECT().Delete(class.ID).Return(nil)

		err := useCase.Delete(class.ID)

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
	useCase := NewClassUseCase(nil, mockRepo, logger)

	class := &models.Class{
		Title:       "title",
		Description: "description",
	}

	fixedClass := &models.Class{
		Title:       "title updated",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mockRepo.EXPECT().GetById(class.ID).Return(class, nil)
		mockRepo.EXPECT().Update(class.ID, fixedClass).Return(fixedClass, nil)

		updatedClass, err := useCase.Update(class.ID, fixedClass)

		assert.NoError(t, err)
		assert.NotNil(t, updatedClass)
		assert.Equal(t, updatedClass, fixedClass)
	})

	t.Run("Update Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		updatedClass, err := useCase.Update(uint(10), fixedClass)

		assert.Error(t, err)
		assert.Nil(t, updatedClass)
		assert.EqualError(t, err, "Not found")
	})
}