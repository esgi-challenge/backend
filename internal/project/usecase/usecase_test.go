package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/project/mock"
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
	useCase := NewProjectUseCase(nil, mockRepo, logger)

	project := &models.Project{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().Create(project).Return(project, nil)

	createdProject, err := useCase.Create(project)

	assert.NoError(t, err)
	assert.NotNil(t, createdProject)
}

func TestGetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewProjectUseCase(nil, mockRepo, logger)

	project := &models.Project{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().GetById(project.ID).Return(project, nil)

	dbProject, err := useCase.GetById(project.ID)

	assert.NoError(t, err)
	assert.NotNil(t, dbProject)
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewProjectUseCase(nil, mockRepo, logger)

	projects := &[]models.Project{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	mockRepo.EXPECT().GetAll().Return(projects, nil)

	dbProjects, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, dbProjects)
	assert.Len(t, *dbProjects, 2)
	assert.Equal(t, &(*projects)[0], &(*dbProjects)[0])
	assert.Equal(t, &(*projects)[1], &(*dbProjects)[1])
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewProjectUseCase(nil, mockRepo, logger)

	project := &models.Project{
		Title:       "title",
		Description: "description",
	}

	t.Run("Delete", func(t *testing.T) {
		mockRepo.EXPECT().GetById(project.ID).Return(project, nil)
		mockRepo.EXPECT().Delete(project.ID).Return(nil)

		err := useCase.Delete(project.ID)

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
	useCase := NewProjectUseCase(nil, mockRepo, logger)

	project := &models.Project{
		Title:       "title",
		Description: "description",
	}

	fixedProject := &models.Project{
		Title:       "title updated",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mockRepo.EXPECT().GetById(project.ID).Return(project, nil)
		mockRepo.EXPECT().Update(project.ID, fixedProject).Return(fixedProject, nil)

		updatedProject, err := useCase.Update(project.ID, fixedProject)

		assert.NoError(t, err)
		assert.NotNil(t, updatedProject)
		assert.Equal(t, updatedProject, fixedProject)
	})

	t.Run("Update Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		updatedProject, err := useCase.Update(uint(10), fixedProject)

		assert.Error(t, err)
		assert.Nil(t, updatedProject)
		assert.EqualError(t, err, "Not found")
	})
}
