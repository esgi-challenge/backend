package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/user/mock"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAll(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()
	logger.InitLogger()
	mockUseCase := mock.NewMockUseCase(ctrl)
	handlers := NewUserHandlers(mockUseCase, nil, logger)

	users := &[]models.User{
		{
			Firstname: "firstname",
			Lastname:  "lastname",
			Email:     "email@gmail.com",
			Password:  "password",
		},
		{
			Firstname: "firstname2",
			Lastname:  "lastname2",
			Email:     "email2@gmail.com",
			Password:  "password",
		},
	}

	t.Run("Get all request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/users/", nil)
		res := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(res)
		ctx.Request = req

		mockUseCase.EXPECT().GetAll().Return(users, nil)

		handlerFunc := handlers.GetAll()
		handlerFunc(ctx)

		assert.Equal(t, http.StatusOK, res.Code)
	})
}
