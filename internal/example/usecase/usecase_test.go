package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/example/mock"
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
  useCase := NewExampleUseCase(mockRepo, nil, logger)

  example := &models.Example{
    Title: "title",
    Description: "description",
  }

  mockRepo.EXPECT().Create(example).Return(example, nil)

  createdExample, err := useCase.Create(example)

  assert.NoError(t, err)
  assert.NotNil(t, createdExample)
}

func TestGetById(t *testing.T) {
  t.Parallel()

  ctrl := gomock.NewController(t)
  defer ctrl.Finish()

  logger := logger.NewLogger()
  mockRepo := mock.NewMockRepository(ctrl)
  useCase := NewExampleUseCase(mockRepo, nil, logger)

  example := &models.Example{
    Title: "title",
    Description: "description",
  }

  mockRepo.EXPECT().GetById(example.ID).Return(example, nil)

  dbExample, err := useCase.GetById(example.ID)

  assert.NoError(t, err)
  assert.NotNil(t, dbExample)
}

func TestGetAll(t *testing.T) {
  t.Parallel()

  ctrl := gomock.NewController(t)
  defer ctrl.Finish()

  logger := logger.NewLogger()
  mockRepo := mock.NewMockRepository(ctrl)
  useCase := NewExampleUseCase(mockRepo, nil, logger)

  examples := &[]models.Example{
    {
      Title: "title1",
      Description: "description1",
    },
    {
      Title: "title2",
      Description: "description2",
    },
  }

  mockRepo.EXPECT().GetAll().Return(examples, nil)

  dbExamples, err := useCase.GetAll()

  assert.NoError(t, err)
  assert.NotNil(t, dbExamples)
  assert.Len(t, *dbExamples, 2)
  assert.Equal(t, &(*examples)[0], &(*dbExamples)[0])
  assert.Equal(t, &(*examples)[1], &(*dbExamples)[1])
}


func TestDelete(t *testing.T) {
  t.Parallel()

  ctrl := gomock.NewController(t)
  defer ctrl.Finish()

  logger := logger.NewLogger()
  mockRepo := mock.NewMockRepository(ctrl)
  useCase := NewExampleUseCase(mockRepo, nil, logger)

  example := &models.Example{
    Title: "title",
    Description: "description",
  }

  t.Run("Delete", func(t *testing.T) {
    mockRepo.EXPECT().GetById(example.ID).Return(example, nil)
    mockRepo.EXPECT().Delete(example.ID).Return(nil)

    err := useCase.Delete(example.ID)

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
  useCase := NewExampleUseCase(mockRepo, nil, logger)

  example := &models.Example{
    Title: "title",
    Description: "description",
  }

  fixedExample := &models.Example{
    Title: "title updated",
    Description: "description",
  }

  t.Run("Update", func(t *testing.T) {
    mockRepo.EXPECT().GetById(example.ID).Return(example, nil)
    mockRepo.EXPECT().Update(example.ID, fixedExample).Return(fixedExample, nil)

    updatedExample, err := useCase.Update(example.ID, fixedExample)

    assert.NoError(t, err)
    assert.NotNil(t, updatedExample)
    assert.Equal(t, updatedExample, fixedExample)
  })

  t.Run("Update Not found", func(t *testing.T) {
    mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

    updatedExample, err := useCase.Update(uint(10), fixedExample)

    assert.Error(t, err)
    assert.Nil(t, updatedExample)
    assert.EqualError(t, err, "Not found")
  })
}
