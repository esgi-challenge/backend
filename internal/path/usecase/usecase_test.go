package usecase

import (
	"testing"
	"errors"

	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/path/mock"
	schoolMock "github.com/esgi-challenge/backend/internal/school/mock"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"go.uber.org/mock/gomock"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestCreatePath(t *testing.T) {
  t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPathRepo := mock.NewMockRepository(ctrl)
	mockSchoolRepo := schoolMock.NewMockRepository(ctrl)
	logger := logger.NewLogger()

	useCase := NewPathUseCase(nil, mockPathRepo, mockSchoolRepo, logger)

	t.Run("success", func(t *testing.T) {
		user := &models.User{GormModel: models.GormModel{ID: 1}}
		school := &models.School{GormModel: models.GormModel{ID: 1}, UserID: 1}
		path := &models.Path{SchoolId: 1}

		mockSchoolRepo.EXPECT().GetById(uint(1)).Return(school, nil)
		mockPathRepo.EXPECT().Create(path).Return(path, nil)

		createdPath, err := useCase.Create(user, path)
		assert.NoError(t, err)
		assert.Equal(t, path, createdPath)
	})

	t.Run("school not owned by user", func(t *testing.T) {
		user := &models.User{GormModel: models.GormModel{ID: 1}}
		school := &models.School{GormModel: models.GormModel{ID: 1}, UserID: 2}
		path := &models.Path{SchoolId: 1}

		mockSchoolRepo.EXPECT().GetById(uint(1)).Return(school, nil)

		createdPath, err := useCase.Create(user, path)
		assert.Error(t, err)
		assert.Nil(t, createdPath)
		assert.Equal(t, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "This school is not yours",
		}, err)
	})

	t.Run("school not found", func(t *testing.T) {
		user := &models.User{GormModel: models.GormModel{ID: 1}}
		path := &models.Path{SchoolId: 1}

		mockSchoolRepo.EXPECT().GetById(uint(1)).Return(nil, errors.New("not found"))

		createdPath, err := useCase.Create(user, path)
		assert.Error(t, err)
		assert.Nil(t, createdPath)
	})
}

func TestGetAllPaths(t *testing.T) {
  t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPathRepo := mock.NewMockRepository(ctrl)
	mockSchoolRepo := schoolMock.NewMockRepository(ctrl)
	logger := logger.NewLogger()

	useCase := NewPathUseCase(nil, mockPathRepo, mockSchoolRepo, logger)

	t.Run("success", func(t *testing.T) {
		paths := &[]models.Path{{GormModel: models.GormModel{ID: 1}}, {GormModel: models.GormModel{ID: 2}}}
		mockPathRepo.EXPECT().GetAll().Return(paths, nil)

		result, err := useCase.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, paths, result)
	})

	t.Run("error", func(t *testing.T) {
		mockPathRepo.EXPECT().GetAll().Return(nil, errors.New("some error"))

		result, err := useCase.GetAll()
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestUpdatePath(t *testing.T) {
  t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPathRepo := mock.NewMockRepository(ctrl)
	mockSchoolRepo := schoolMock.NewMockRepository(ctrl)
	logger := logger.NewLogger()

	useCase := NewPathUseCase(nil, mockPathRepo, mockSchoolRepo, logger)

	t.Run("success", func(t *testing.T) {
		path := &models.Path{GormModel: models.GormModel{ID: 1}}
		updatedPath := &models.Path{ShortName: "Updated"}
		mockPathRepo.EXPECT().GetById(uint(1)).Return(path, nil)
		mockPathRepo.EXPECT().Update(uint(1), updatedPath).Return(updatedPath, nil)

		result, err := useCase.Update(uint(1), updatedPath)
		assert.NoError(t, err)
		assert.Equal(t, updatedPath, result)
	})

	t.Run("error", func(t *testing.T) {
		updatedPath := &models.Path{ShortName: "Updated"}
		mockPathRepo.EXPECT().GetById(uint(1)).Return(nil, errors.New("not found"))

		result, err := useCase.Update(uint(1), updatedPath)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
