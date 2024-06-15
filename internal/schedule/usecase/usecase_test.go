package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/schedule/mock"
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
	useCase := NewScheduleUseCase(nil, mockRepo, logger)

	schedule := &models.Schedule{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().Create(schedule).Return(schedule, nil)

	createdSchedule, err := useCase.Create(schedule)

	assert.NoError(t, err)
	assert.NotNil(t, createdSchedule)
}

func TestGetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewScheduleUseCase(nil, mockRepo, logger)

	schedule := &models.Schedule{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().GetById(schedule.ID).Return(schedule, nil)

	dbSchedule, err := useCase.GetById(schedule.ID)

	assert.NoError(t, err)
	assert.NotNil(t, dbSchedule)
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewScheduleUseCase(nil, mockRepo, logger)

	schedules := &[]models.Schedule{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	mockRepo.EXPECT().GetAll().Return(schedules, nil)

	dbSchedules, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, dbSchedules)
	assert.Len(t, *dbSchedules, 2)
	assert.Equal(t, &(*schedules)[0], &(*dbSchedules)[0])
	assert.Equal(t, &(*schedules)[1], &(*dbSchedules)[1])
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewScheduleUseCase(nil, mockRepo, logger)

	schedule := &models.Schedule{
		Title:       "title",
		Description: "description",
	}

	t.Run("Delete", func(t *testing.T) {
		mockRepo.EXPECT().GetById(schedule.ID).Return(schedule, nil)
		mockRepo.EXPECT().Delete(schedule.ID).Return(nil)

		err := useCase.Delete(schedule.ID)

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
	useCase := NewScheduleUseCase(nil, mockRepo, logger)

	schedule := &models.Schedule{
		Title:       "title",
		Description: "description",
	}

	fixedSchedule := &models.Schedule{
		Title:       "title updated",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mockRepo.EXPECT().GetById(schedule.ID).Return(schedule, nil)
		mockRepo.EXPECT().Update(schedule.ID, fixedSchedule).Return(fixedSchedule, nil)

		updatedSchedule, err := useCase.Update(schedule.ID, fixedSchedule)

		assert.NoError(t, err)
		assert.NotNil(t, updatedSchedule)
		assert.Equal(t, updatedSchedule, fixedSchedule)
	})

	t.Run("Update Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		updatedSchedule, err := useCase.Update(uint(10), fixedSchedule)

		assert.Error(t, err)
		assert.Nil(t, updatedSchedule)
		assert.EqualError(t, err, "Not found")
	})
}