package request

import (
	"errors"
	"strings"

	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateJSON[T interface{}](input T, ctx *gin.Context) (T, error) {
	bindError := ctx.ShouldBindJSON(&input)
	if bindError != nil {
		var missingFields []string
		for _, err := range bindError.(validator.ValidationErrors) {
			missingFields = append(missingFields, err.Field())
		}

		return input, errors.New("Missing fields: " + strings.ToLower(strings.Join(missingFields, ", ")))
	}

	validate := validator.New()
	validateError := validate.Struct(input)
	if validateError != nil {
		var invalidFields []string
		for _, err := range validateError.(validator.ValidationErrors) {
			invalidFields = append(invalidFields, err.Field())
		}

		return input, errors.New("Invalid fields: " + strings.ToLower(strings.Join(invalidFields, ", ")))
	}

	return input, nil
}

func ValidateRole(secretKey string, ctx *gin.Context, role models.UserKind) (*models.User, error) {
	token := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]

	user, err := jwt.DecryptToken(secretKey, token)

	if err != nil {
		return nil, err
	}

	if user.UserKind < role {
		return nil, errors.New("forbidden")
	}

	return user, nil
}
