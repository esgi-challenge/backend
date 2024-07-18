package http

import (
	"bytes"
	"encoding/json"
	"errors"

	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/project/mock"
	"github.com/esgi-challenge/backend/pkg/jwt"
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
	cfg := &config.Config{JwtSecret: "secret"}
	handlers := NewProjectHandlers(cfg, mockUseCase, logger)

	id := uint(1)
	project := &models.ProjectCreate{
		Title:      "title",
		EndDate:    "10/10/10/",
		CourseId:   &id,
		ClassId:    &id,
		DocumentId: &id,
	}

	body, err := json.Marshal(project)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	adminUser := &models.User{
		UserKind: models.NewUserKind(models.ADMINISTRATOR),
	}
	token, _ := jwt.Generate(cfg.JwtSecret, adminUser)

	t.Run("Create request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/projects", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		expectedProject := &models.Project{
			Title:   "title",
			EndDate: "10/10/10/",
		}

		mockUseCase.EXPECT().Create(gomock.Any(), gomock.Any()).Return(expectedProject, nil)

		handler := handlers.Create()
		handler(ctx)

		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("Create request server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/paths", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		mockUseCase.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("random server error"))

		handler := handlers.Create()
		handler(ctx)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
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
	cfg := &config.Config{JwtSecret: "secret"}
	handlers := NewProjectHandlers(cfg, mockUseCase, logger)

	paths := &[]models.Project{
		{
			Title:   "title1",
			EndDate: "10/10/10",
		},
		{
			Title:   "title2",
			EndDate: "20/20/20",
		},
	}

	adminUser := &models.User{
		UserKind: models.NewUserKind(models.ADMINISTRATOR),
	}
	token, _ := jwt.Generate(cfg.JwtSecret, adminUser)

	t.Run("Get all request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/paths", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Request = req

		mockUseCase.EXPECT().GetAll(gomock.Any()).Return(paths, nil)

		handlerFunc := handlers.GetAll()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Get all request server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/paths", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Request = req

		mockUseCase.EXPECT().GetAll(gomock.Any()).Return(nil, errors.New("random server error"))

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
	cfg := &config.Config{JwtSecret: "secret"}
	handlers := NewProjectHandlers(cfg, mockUseCase, logger)

	id := uint(1)
	path := &models.ProjectUpdate{
		Title:      "name",
		EndDate:    "10/10/10",
		CourseId:   &id,
		ClassId:    &id,
		DocumentId: &id,
	}

	adminUser := &models.User{
		UserKind: models.NewUserKind(models.ADMINISTRATOR),
	}
	token, _ := jwt.Generate(cfg.JwtSecret, adminUser)

	body, err := json.Marshal(path)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	t.Run("Update request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/projects/1", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}
		ctx.Request = req

		expectedProject := &models.Project{
			Title:   "updated",
			EndDate: "20/20/20",
		}

		mockUseCase.EXPECT().Update(adminUser, uint(1), gomock.Any()).Return(expectedProject, nil)

		handler := handlers.Update()
		handler(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Update request server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/projects/1", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Request = req

		mockUseCase.EXPECT().Update(adminUser, uint(1), gomock.Any()).Return(nil, errors.New("random server error"))

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
	cfg := &config.Config{JwtSecret: "secret"}
	handlers := NewProjectHandlers(cfg, mockUseCase, logger)

	adminUser := &models.User{
		UserKind: models.NewUserKind(models.ADMINISTRATOR),
	}
	token, _ := jwt.Generate(cfg.JwtSecret, adminUser)

	t.Run("Delete request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/projects/1", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Request = req

		mockUseCase.EXPECT().Delete(adminUser, uint(1)).Return(nil)

		handlerFunc := handlers.Delete()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Delete request not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/projects/10", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "10"}}
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Request = req

		mockUseCase.EXPECT().Delete(adminUser, uint(10)).Return(gorm.ErrRecordNotFound)

		handlerFunc := handlers.Delete()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}
