package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/school/mock"
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
	useCase := NewSchoolUseCase(nil, mockRepo, logger)

	school := &models.School{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().Create(school).Return(school, nil)

	createdSchool, err := useCase.Create(school)

	assert.NoError(t, err)
	assert.NotNil(t, createdSchool)
}

func TestGetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewSchoolUseCase(nil, mockRepo, logger)

	school := &models.School{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().GetById(school.ID).Return(school, nil)

	dbSchool, err := useCase.GetById(school.ID)

	assert.NoError(t, err)
	assert.NotNil(t, dbSchool)
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewSchoolUseCase(nil, mockRepo, logger)

	schools := &[]models.School{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	mockRepo.EXPECT().GetAll().Return(schools, nil)

	dbSchools, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, dbSchools)
	assert.Len(t, *dbSchools, 2)
	assert.Equal(t, &(*schools)[0], &(*dbSchools)[0])
	assert.Equal(t, &(*schools)[1], &(*dbSchools)[1])
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewSchoolUseCase(nil, mockRepo, logger)

	school := &models.School{
		Title:       "title",
		Description: "description",
	}

	t.Run("Delete", func(t *testing.T) {
		mockRepo.EXPECT().GetById(school.ID).Return(school, nil)
		mockRepo.EXPECT().Delete(school.ID).Return(nil)

		err := useCase.Delete(school.ID)

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
	useCase := NewSchoolUseCase(nil, mockRepo, logger)

	school := &models.School{
		Title:       "title",
		Description: "description",
	}

	fixedSchool := &models.School{
		Title:       "title updated",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mockRepo.EXPECT().GetById(school.ID).Return(school, nil)
		mockRepo.EXPECT().Update(school.ID, fixedSchool).Return(fixedSchool, nil)

		updatedSchool, err := useCase.Update(school.ID, fixedSchool)

		assert.NoError(t, err)
		assert.NotNil(t, updatedSchool)
		assert.Equal(t, updatedSchool, fixedSchool)
	})

	t.Run("Update Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		updatedSchool, err := useCase.Update(uint(10), fixedSchool)

		assert.Error(t, err)
		assert.Nil(t, updatedSchool)
		assert.EqualError(t, err, "Not found")
	})
}
