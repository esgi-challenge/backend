package usecase

import (
	"errors"
	"testing"

	classMock "github.com/esgi-challenge/backend/internal/class/mock"
	courseMock "github.com/esgi-challenge/backend/internal/course/mock"
	documentMock "github.com/esgi-challenge/backend/internal/document/mock"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/project/mock"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreatePath(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProjectRepo := mock.NewMockRepository(ctrl)
	mockCourseUsecase := courseMock.NewMockUseCase(ctrl)
	mockClassUsecase := classMock.NewMockUseCase(ctrl)
	mockDocumentUsecase := documentMock.NewMockUseCase(ctrl)
	logger := logger.NewLogger()

	useCase := NewProjectUseCase(nil, mockProjectRepo, mockCourseUsecase, mockClassUsecase, mockDocumentUsecase, logger)

	t.Run("success", func(t *testing.T) {
		user := &models.User{GormModel: models.GormModel{ID: 1}}
		project := &models.Project{Title: "title", EndDate: "10/10/10"}

		mockCourseUsecase.EXPECT().GetById(gomock.Any()).Return(nil, nil)
		mockDocumentUsecase.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, nil)
		mockClassUsecase.EXPECT().GetById(gomock.Any()).Return(nil, nil)
		mockProjectRepo.EXPECT().Create(project).Return(project, nil)
		mockProjectRepo.EXPECT().GetPreloadById(project.ID).Return(project, nil)

		createdProject, err := useCase.Create(user, project)
		assert.NoError(t, err)
		assert.Equal(t, project, createdProject)
	})
}

func TestGetAllPaths(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProjectRepo := mock.NewMockRepository(ctrl)
	mockCourseUsecase := courseMock.NewMockUseCase(ctrl)
	mockClassUsecase := classMock.NewMockUseCase(ctrl)
	mockDocumentUsecase := documentMock.NewMockUseCase(ctrl)
	logger := logger.NewLogger()

	useCase := NewProjectUseCase(nil, mockProjectRepo, mockCourseUsecase, mockClassUsecase, mockDocumentUsecase, logger)

	user := &models.User{
		UserKind: models.NewUserKind(models.TEACHER),
	}

	t.Run("success", func(t *testing.T) {
		projects := &[]models.Project{{GormModel: models.GormModel{ID: 1}}, {GormModel: models.GormModel{ID: 2}}}
		mockProjectRepo.EXPECT().GetAllByTeacher(user).Return(projects, nil)

		result, err := useCase.GetAll(user)
		assert.NoError(t, err)
		assert.Equal(t, projects, result)
	})

	t.Run("error", func(t *testing.T) {
		mockProjectRepo.EXPECT().GetAllByTeacher(user).Return(nil, errors.New("some error"))

		result, err := useCase.GetAll(user)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
