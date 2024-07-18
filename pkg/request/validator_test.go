package request

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidateJSON(t *testing.T) {
	t.Parallel()

	type TestStruct struct {
		Name string `json:"name" validate:"required"`
		Age  int    `json:"age" validate:"required"`
	}

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/test", func(ctx *gin.Context) {
		var input TestStruct
		validatedInput, err := ValidateJSON(input, ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, validatedInput)
	})

	t.Run("valid input", func(t *testing.T) {
		body := `{"name": "John", "age": 30}`
		req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "John")
		assert.Contains(t, resp.Body.String(), "30")
	})
}

func TestValidateRole(t *testing.T) {
	t.Parallel()

	secretKey := "secret"
	user := &models.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		UserKind:  new(models.UserKind),
	}
	*user.UserKind = models.ADMINISTRATOR

	token, _ := jwt.Generate(secretKey, user)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test", func(ctx *gin.Context) {
		validatedUser, err := ValidateRole(secretKey, ctx, models.TEACHER)
		if err != nil {
			ctx.JSON(err.(errorHandler.HttpError).HttpStatus, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, validatedUser)
	})

	t.Run("missing authorization header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.Contains(t, resp.Body.String(), "Unauthorized")
	})

	t.Run("invalid token format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.Contains(t, resp.Body.String(), "Unauthorized")
	})

	t.Run("valid token with sufficient role", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "John")
		assert.Contains(t, resp.Body.String(), "Doe")
	})
}

func TestValidateRoleWithoutHeader(t *testing.T) {
	t.Parallel()

	secretKey := "secret"
	user := &models.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		UserKind:  new(models.UserKind),
	}
	*user.UserKind = models.ADMINISTRATOR

	token, _ := jwt.Generate(secretKey, user)

	t.Run("valid token with sufficient role", func(t *testing.T) {
		validatedUser, err := ValidateRoleWithoutHeader(secretKey, token, models.TEACHER)
		assert.NoError(t, err)
		assert.NotNil(t, validatedUser)
		assert.Equal(t, "John", validatedUser.Firstname)
		assert.Equal(t, "Doe", validatedUser.Lastname)
	})

	t.Run("valid token with insufficient role", func(t *testing.T) {
		*user.UserKind = models.STUDENT
		token, _ := jwt.Generate(secretKey, user)
		validatedUser, err := ValidateRoleWithoutHeader(secretKey, token, models.TEACHER)
		assert.Error(t, err)
		assert.Nil(t, validatedUser)
		assert.Equal(t, errorHandler.Forbidden.Error(), err.Error())
	})
}
