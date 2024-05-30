package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/esgi-challenge/backend/internal/path/mock"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	logger.InitLogger()
	mockUseCase := mock.NewMockUseCase(ctrl)
	handlers := NewPathHandlers(nil, mockUseCase, logger)

	path := &models.PathCreate{
		Title:       "longtitle",
		Description: "description",
	}

	body, err := json.Marshal(path)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	t.Run("Create request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/paths", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		expectedPath := &models.Path{
			Title:       "title",
			Description: "description",
		}
		mockUseCase.EXPECT().Create(gomock.Any()).Return(expectedPath, nil)

		handler := handlers.Create()
		handler(ctx)

		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("Create request bad body", func(t *testing.T) {
		badPath := &models.PathCreate{
			Title:       "bad",
			Description: "description",
		}

		body, err := json.Marshal(badPath)
		if err != nil {
			t.Fatalf("Failed to marshal JSON: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/paths", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		handler := handlers.Create()
		handler(ctx)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Create request server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/paths", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		mockUseCase.EXPECT().Create(gomock.Any()).Return(nil, errors.New("random server error"))

		handler := handlers.Create()
		handler(ctx)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestGetById(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	logger.InitLogger()
	mockUseCase := mock.NewMockUseCase(ctrl)
	handlers := NewPathHandlers(nil, mockUseCase, logger)

	path := &models.Path{
		Title:       "title",
		Description: "description",
	}

	t.Run("Get by id request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/paths/"+strconv.Itoa(int(path.ID)), nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: strconv.Itoa(int(path.ID))}}
		ctx.Request = req

		mockUseCase.EXPECT().GetById(path.ID).Return(path, nil)

		handlerFunc := handlers.GetById()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Get by id request wrong url param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/paths/badParam", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "badParam"}}
		ctx.Request = req

		handlerFunc := handlers.GetById()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Get by id request not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/paths/10", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "10"}}
		ctx.Request = req

		mockUseCase.EXPECT().GetById(uint(10)).Return(nil, gorm.ErrRecordNotFound)

		handlerFunc := handlers.GetById()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}

func TestGetAll(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	logger.InitLogger()
	mockUseCase := mock.NewMockUseCase(ctrl)
	handlers := NewPathHandlers(nil, mockUseCase, logger)

	paths := &[]models.Path{
		{
			Title:       "title1",
			Description: "description1",
		},
		{
			Title:       "title2",
			Description: "description2",
		},
	}

	t.Run("Get all request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/paths", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		mockUseCase.EXPECT().GetAll().Return(paths, nil)

		handlerFunc := handlers.GetAll()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Get all request server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/paths", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		mockUseCase.EXPECT().GetAll().Return(nil, errors.New("random server error"))

		handlerFunc := handlers.GetAll()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestUpdate(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	logger.InitLogger()
	mockUseCase := mock.NewMockUseCase(ctrl)
	handlers := NewPathHandlers(nil, mockUseCase, logger)

	path := &models.PathUpdate{
		Title:       "longtitle",
		Description: "description",
	}

	body, err := json.Marshal(path)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	t.Run("Update request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/paths/1", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}
		ctx.Request = req

		expectedPath := &models.Path{
			Title:       "updated",
			Description: "description",
		}
		mockUseCase.EXPECT().Update(uint(1), gomock.Any()).Return(expectedPath, nil)

		handler := handlers.Update()
		handler(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Update request wrong url param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/paths/badParam", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "badParam"}}
		ctx.Request = req

		handlerFunc := handlers.Update()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Update request bad body", func(t *testing.T) {
		badPath := &models.PathUpdate{
			Title:       "bad",
			Description: "description",
		}

		body, err := json.Marshal(badPath)
		if err != nil {
			t.Fatalf("Failed to marshal JSON: %v", err)
		}

		req := httptest.NewRequest(http.MethodPut, "/api/paths/1", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}
		ctx.Request = req

		handler := handlers.Update()
		handler(ctx)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Update request server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/paths/1", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}
		ctx.Request = req

		mockUseCase.EXPECT().Update(uint(1), gomock.Any()).Return(nil, errors.New("random server error"))

		handler := handlers.Update()
		handler(ctx)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	logger.InitLogger()
	mockUseCase := mock.NewMockUseCase(ctrl)
	handlers := NewPathHandlers(nil, mockUseCase, logger)

	t.Run("Delete request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/paths/1", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}
		ctx.Request = req

		mockUseCase.EXPECT().Delete(uint(1)).Return(nil)

		handlerFunc := handlers.Delete()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Delete request wrong url param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/paths/badParam", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "badParam"}}
		ctx.Request = req

		handlerFunc := handlers.Delete()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Delete request not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/paths/10", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "10"}}
		ctx.Request = req

		mockUseCase.EXPECT().Delete(uint(10)).Return(gorm.ErrRecordNotFound)

		handlerFunc := handlers.Delete()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}