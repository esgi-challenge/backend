package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/esgi-challenge/backend/internal/example/mock"
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
	handlers := NewExampleHandlers(mockUseCase, nil, logger)

	example := &models.ExampleCreate{
		Title:       "longtitle",
		Description: "description",
	}

	body, err := json.Marshal(example)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	t.Run("Create request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/examples", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		expectedExample := &models.Example{
			Title:       "title",
			Description: "description",
		}
		mockUseCase.EXPECT().Create(gomock.Any()).Return(expectedExample, nil)

		handler := handlers.Create()
		handler(ctx)

		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("Create request bad body", func(t *testing.T) {
		badExample := &models.ExampleCreate{
			Title:       "bad",
			Description: "description",
		}

		body, err := json.Marshal(badExample)
		if err != nil {
			t.Fatalf("Failed to marshal JSON: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/examples", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		handler := handlers.Create()
		handler(ctx)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Create request server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/examples", bytes.NewBuffer(body))
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
	handlers := NewExampleHandlers(mockUseCase, nil, logger)

	example := &models.Example{
		Title:       "title",
		Description: "description",
	}

	t.Run("Get by id request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/examples/"+strconv.Itoa(int(example.ID)), nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: strconv.Itoa(int(example.ID))}}
		ctx.Request = req

		mockUseCase.EXPECT().GetById(example.ID).Return(example, nil)

		handlerFunc := handlers.GetById()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Get by id request wrong url param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/examples/badParam", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "badParam"}}
		ctx.Request = req

		handlerFunc := handlers.GetById()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Get by id request not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/examples/10", nil)
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
	handlers := NewExampleHandlers(mockUseCase, nil, logger)

	examples := &[]models.Example{
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
		req := httptest.NewRequest(http.MethodGet, "/api/examples", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		mockUseCase.EXPECT().GetAll().Return(examples, nil)

		handlerFunc := handlers.GetAll()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Get all request server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/examples", nil)
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
	handlers := NewExampleHandlers(mockUseCase, nil, logger)

	example := &models.ExampleUpdate{
		Title:       "longtitle",
		Description: "description",
	}

	body, err := json.Marshal(example)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	t.Run("Update request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/examples/1", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}
		ctx.Request = req

		expectedExample := &models.Example{
			Title:       "updated",
			Description: "description",
		}
		mockUseCase.EXPECT().Update(uint(1), gomock.Any()).Return(expectedExample, nil)

		handler := handlers.Update()
		handler(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Update request wrong url param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/examples/badParam", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "badParam"}}
		ctx.Request = req

		handlerFunc := handlers.Update()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Update request bad body", func(t *testing.T) {
		badExample := &models.ExampleUpdate{
			Title:       "bad",
			Description: "description",
		}

		body, err := json.Marshal(badExample)
		if err != nil {
			t.Fatalf("Failed to marshal JSON: %v", err)
		}

		req := httptest.NewRequest(http.MethodPut, "/api/examples/1", bytes.NewBuffer(body))
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
		req := httptest.NewRequest(http.MethodPut, "/api/examples/1", bytes.NewBuffer(body))
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
	handlers := NewExampleHandlers(mockUseCase, nil, logger)

	t.Run("Delete request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/examples/1", nil)
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
		req := httptest.NewRequest(http.MethodDelete, "/api/examples/badParam", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "badParam"}}
		ctx.Request = req

		handlerFunc := handlers.Delete()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Delete request not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/examples/10", nil)
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
