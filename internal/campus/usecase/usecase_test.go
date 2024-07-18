package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/campus/mock"
	schoolMock "github.com/esgi-challenge/backend/internal/school/mock"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
)

func TestCreateCampus(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCampusRepo := mock.NewMockRepository(ctrl)
	mockSchoolRepo := schoolMock.NewMockRepository(ctrl)
	logger := logger.NewLogger()

	useCase := NewCampusUseCase(nil, mockCampusRepo, mockSchoolRepo, logger)

	t.Run("success", func(t *testing.T) {
		user := &models.User{GormModel: models.GormModel{ID: 1}}
		school := &models.School{GormModel: models.GormModel{ID: 1}, UserID: 1}
		campus := &models.Campus{SchoolId: 1}

		mockSchoolRepo.EXPECT().GetById(uint(1)).Return(school, nil)
		mockCampusRepo.EXPECT().Create(campus).Return(campus, nil)

		createdCampus, err := useCase.Create(user, campus)
		assert.NoError(t, err)
		assert.Equal(t, campus, createdCampus)
	})

	t.Run("school not owned by user", func(t *testing.T) {
		user := &models.User{GormModel: models.GormModel{ID: 1}}
		school := &models.School{GormModel: models.GormModel{ID: 1}, UserID: 2}
		campus := &models.Campus{SchoolId: 1}

		mockSchoolRepo.EXPECT().GetById(uint(1)).Return(school, nil)

		createdCampus, err := useCase.Create(user, campus)
		assert.Error(t, err)
		assert.Nil(t, createdCampus)
		assert.Equal(t, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "This school is not yours",
		}, err)
	})

	t.Run("school not found", func(t *testing.T) {
		user := &models.User{GormModel: models.GormModel{ID: 1}}
		campus := &models.Campus{SchoolId: 1}

		mockSchoolRepo.EXPECT().GetById(uint(1)).Return(nil, errors.New("not found"))

		createdCampus, err := useCase.Create(user, campus)
		assert.Error(t, err)
		assert.Nil(t, createdCampus)
	})
}

func TestGetAllCampus(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCampusRepo := mock.NewMockRepository(ctrl)
	mockSchoolRepo := schoolMock.NewMockRepository(ctrl)
	logger := logger.NewLogger()

	useCase := NewCampusUseCase(nil, mockCampusRepo, mockSchoolRepo, logger)

	t.Run("success", func(t *testing.T) {
		campus := &[]models.Campus{{GormModel: models.GormModel{ID: 1}}, {GormModel: models.GormModel{ID: 2}}}
		mockCampusRepo.EXPECT().GetAll().Return(campus, nil)

		result, err := useCase.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, campus, result)
	})

	t.Run("error", func(t *testing.T) {
		mockCampusRepo.EXPECT().GetAll().Return(nil, errors.New("some error"))

		result, err := useCase.GetAll()
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
