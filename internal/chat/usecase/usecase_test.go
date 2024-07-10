package usecase

import (
	"errors"
	"testing"

	"github.com/esgi-challenge/backend/internal/chat/mock"
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
	useCase := NewChatUseCase(nil, mockRepo, logger)

	chat := &models.Chat{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().Create(chat).Return(chat, nil)

	createdChat, err := useCase.Create(chat)

	assert.NoError(t, err)
	assert.NotNil(t, createdChat)
}

func TestGetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewChatUseCase(nil, mockRepo, logger)

	chat := &models.Chat{
		Title:       "title",
		Description: "description",
	}

	mockRepo.EXPECT().GetById(chat.ID).Return(chat, nil)

	dbChat, err := useCase.GetById(chat.ID)

	assert.NoError(t, err)
	assert.NotNil(t, dbChat)
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewChatUseCase(nil, mockRepo, logger)

	chats := &[]models.Chat{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	mockRepo.EXPECT().GetAll().Return(chats, nil)

	dbChats, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, dbChats)
	assert.Len(t, *dbChats, 2)
	assert.Equal(t, &(*chats)[0], &(*dbChats)[0])
	assert.Equal(t, &(*chats)[1], &(*dbChats)[1])
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	mockRepo := mock.NewMockRepository(ctrl)
	useCase := NewChatUseCase(nil, mockRepo, logger)

	chat := &models.Chat{
		Title:       "title",
		Description: "description",
	}

	t.Run("Delete", func(t *testing.T) {
		mockRepo.EXPECT().GetById(chat.ID).Return(chat, nil)
		mockRepo.EXPECT().Delete(chat.ID).Return(nil)

		err := useCase.Delete(chat.ID)

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
	useCase := NewChatUseCase(nil, mockRepo, logger)

	chat := &models.Chat{
		Title:       "title",
		Description: "description",
	}

	fixedChat := &models.Chat{
		Title:       "title updated",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mockRepo.EXPECT().GetById(chat.ID).Return(chat, nil)
		mockRepo.EXPECT().Update(chat.ID, fixedChat).Return(fixedChat, nil)

		updatedChat, err := useCase.Update(chat.ID, fixedChat)

		assert.NoError(t, err)
		assert.NotNil(t, updatedChat)
		assert.Equal(t, updatedChat, fixedChat)
	})

	t.Run("Update Not found", func(t *testing.T) {
		mockRepo.EXPECT().GetById(uint(10)).Return(nil, errors.New("Not found"))

		updatedChat, err := useCase.Update(uint(10), fixedChat)

		assert.Error(t, err)
		assert.Nil(t, updatedChat)
		assert.EqualError(t, err, "Not found")
	})
}
