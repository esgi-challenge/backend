package http

import (
	"bytes"
	"encoding/json"
	"errors"

	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/campus/mock"
	"github.com/esgi-challenge/backend/internal/models"
	schoolMock "github.com/esgi-challenge/backend/internal/school/mock"
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
	mockSchoolUseCase := schoolMock.NewMockUseCase(ctrl)
	cfg := &config.Config{JwtSecret: "secret"}
	handlers := NewCampusHandlers(cfg, mockUseCase, mockSchoolUseCase, logger, nil)

	campus := &models.CampusCreate{
		Name: "name",
		Location:  "location",
    Latitude: 1,
    Longitude: 1,
	}

	body, err := json.Marshal(campus)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	adminUser := &models.User{
		UserKind: models.NewUserKind(models.ADMINISTRATOR),
	}
	token, _ := jwt.Generate(cfg.JwtSecret, adminUser)

	t.Run("Create request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/campus", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		expectedCampus := &models.Campus{
			Name: "name",
			Location:  "location",
      Latitude: 1,
      Longitude: 1,
		}

		mockSchoolUseCase.EXPECT().GetByUser(gomock.Any()).Return(&models.School{GormModel: models.GormModel{ID: 1}}, nil)
		mockUseCase.EXPECT().Create(gomock.Any(), gomock.Any()).Return(expectedCampus, nil)

		handler := handlers.Create()
		handler(ctx)

		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("Create request server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/campus", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		mockSchoolUseCase.EXPECT().GetByUser(gomock.Any()).Return(&models.School{GormModel: models.GormModel{ID: 1}}, nil)
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
	mockSchoolUseCase := schoolMock.NewMockUseCase(ctrl)
	cfg := &config.Config{JwtSecret: "secret"}
	handlers := NewCampusHandlers(cfg, mockUseCase, mockSchoolUseCase, logger, nil)

	campuses := &[]models.Campus{
		{
			Name: "name1",
			Location:  "location1",
      Latitude: 1,
      Longitude: 1,
		},
		{
			Name: "name2",
			Location:  "location2",
      Latitude: 2,
      Longitude: 2,
		},
	}

	adminUser := &models.User{
		UserKind: models.NewUserKind(models.ADMINISTRATOR),
	}
	token, _ := jwt.Generate(cfg.JwtSecret, adminUser)

	t.Run("Get all request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/campus", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Request = req

		mockSchoolUseCase.EXPECT().GetByUser(gomock.Any()).Return(&models.School{GormModel: models.GormModel{ID: 1}}, nil)
		mockUseCase.EXPECT().GetAllBySchoolId(gomock.Any()).Return(campuses, nil)

		handlerFunc := handlers.GetAll()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Get all request server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/campus", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Request = req

		mockSchoolUseCase.EXPECT().GetByUser(gomock.Any()).Return(&models.School{GormModel: models.GormModel{ID: 1}}, nil)
		mockUseCase.EXPECT().GetAllBySchoolId(gomock.Any()).Return(nil, errors.New("random server error"))

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
	mockSchoolUseCase := schoolMock.NewMockUseCase(ctrl)
	cfg := &config.Config{JwtSecret: "secret"}
	handlers := NewCampusHandlers(cfg, mockUseCase, mockSchoolUseCase, logger, nil)

	campus := &models.CampusUpdate{
		Name: "name",
		Location:  "location",
    Latitude: 1,
    Longitude: 1,
	}

	adminUser := &models.User{
		UserKind: models.NewUserKind(models.ADMINISTRATOR),
	}
	token, _ := jwt.Generate(cfg.JwtSecret, adminUser)

	body, err := json.Marshal(campus)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	t.Run("Update request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/campus/1", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}
		ctx.Request = req

		expectedCampus := &models.Campus{
			Name: "updated",
			Location:  "updated",
      Latitude: 2,
      Longitude: 2,
			SchoolId:  1,
		}

		mockSchoolUseCase.EXPECT().GetByUser(gomock.Any()).Return(&models.School{GormModel: models.GormModel{ID: 1}}, nil)
		mockUseCase.EXPECT().GetById(gomock.Any()).Return(expectedCampus, nil)
		mockUseCase.EXPECT().Update(adminUser, uint(1), gomock.Any()).Return(expectedCampus, nil)

		handler := handlers.Update()
		handler(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Update request server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/campus/1", bytes.NewBuffer(body))
		req.Header.Set("Content-type", "application/json")
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Request = req

		mockSchoolUseCase.EXPECT().GetByUser(gomock.Any()).Return(&models.School{GormModel: models.GormModel{ID: 1}}, nil)
		mockUseCase.EXPECT().GetById(gomock.Any()).Return(&models.Campus{SchoolId: 1}, nil)
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
	mockSchoolUseCase := schoolMock.NewMockUseCase(ctrl)
	cfg := &config.Config{JwtSecret: "secret"}
	handlers := NewCampusHandlers(cfg, mockUseCase, mockSchoolUseCase, logger, nil)

	adminUser := &models.User{
		UserKind: models.NewUserKind(models.ADMINISTRATOR),
	}
	token, _ := jwt.Generate(cfg.JwtSecret, adminUser)

	t.Run("Delete request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/campus/1", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Request = req

		mockSchoolUseCase.EXPECT().GetByUser(gomock.Any()).Return(&models.School{GormModel: models.GormModel{ID: 1}}, nil)
		mockUseCase.EXPECT().GetById(gomock.Any()).Return(&models.Campus{SchoolId: 1}, nil)
		mockUseCase.EXPECT().Delete(adminUser, uint(1)).Return(nil)

		handlerFunc := handlers.Delete()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Delete request not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/campus/10", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Params = gin.Params{{Key: "id", Value: "10"}}
		req.Header.Set("Authorization", "Bearer "+token)
		ctx.Request = req

		mockSchoolUseCase.EXPECT().GetByUser(gomock.Any()).Return(&models.School{GormModel: models.GormModel{ID: 1}}, nil)
		mockUseCase.EXPECT().GetById(gomock.Any()).Return(&models.Campus{SchoolId: 1}, nil)
		mockUseCase.EXPECT().Delete(adminUser, uint(10)).Return(gorm.ErrRecordNotFound)

		handlerFunc := handlers.Delete()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}
