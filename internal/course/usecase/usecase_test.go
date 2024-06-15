package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/course/mock"
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
	useCase := NewCourseUseCase(nil, mockRepo, logger)

	course := &models.Course{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().Create(course).Return(course, nil)

	createdCourse, err := useCase.Create(course)

	assert.NoError(t, err)
	assert.NotNil(t, createdCourse)
}

func TestGetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewCourseUseCase(nil, mockRepo, logger)

	course := &models.Course{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().GetById(course.ID).Return(course, nil)

	dbCourse, err := useCase.GetById(course.ID)

	assert.NoError(t, err)
	assert.NotNil(t, dbCourse)
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewCourseUseCase(nil, mockRepo, logger)

	courses := &[]models.Course{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	mockRepo.EXPECT().GetAll().Return(courses, nil)

	dbCourses, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, dbCourses)
	assert.Len(t, *dbCourses, 2)
	assert.Equal(t, &(*courses)[0], &(*dbCourses)[0])
	assert.Equal(t, &(*courses)[1], &(*dbCourses)[1])
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewCourseUseCase(nil, mockRepo, logger)

	course := &models.Course{
		Title:       "title",
		Description: "description",
	}

	t.Run("Delete", func(t *testing.T) {
		mockRepo.EXPECT().GetById(course.ID).Return(course, nil)
		mockRepo.EXPECT().Delete(course.ID).Return(nil)

		err := useCase.Delete(course.ID)

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
	useCase := NewCourseUseCase(nil, mockRepo, logger)

	course := &models.Course{
		Title:       "title",
		Description: "description",
	}

	fixedCourse := &models.Course{
		Title:       "title updated",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mockRepo.EXPECT().GetById(course.ID).Return(course, nil)
		mockRepo.EXPECT().Update(course.ID, fixedCourse).Return(fixedCourse, nil)

		updatedCourse, err := useCase.Update(course.ID, fixedCourse)

		assert.NoError(t, err)
		assert.NotNil(t, updatedCourse)
		assert.Equal(t, updatedCourse, fixedCourse)
	})

	t.Run("Update Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		updatedCourse, err := useCase.Update(uint(10), fixedCourse)

		assert.Error(t, err)
		assert.Nil(t, updatedCourse)
		assert.EqualError(t, err, "Not found")
	})
}
