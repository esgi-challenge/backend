package request

import (
	"errors"
	"net/http"
	"strings"

	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
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
	if len(ctx.Request.Header["Authorization"]) != 1 {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusUnauthorized,
			HttpError:  errorHandler.Unauthorized.Error(),
		}
	}

	bearer := strings.Split(ctx.Request.Header["Authorization"][0], " ")

	if len(bearer) != 2 {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusUnauthorized,
			HttpError:  errorHandler.Unauthorized.Error(),
		}
	}

	token := bearer[1]

	user, err := jwt.DecryptToken(secretKey, token)

	if err != nil {
		return nil, err
	}

	if user.UserKind < role {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  errorHandler.Forbidden.Error(),
		}
	}

	return user, nil
}

func ValidateRoleWithoutHeader(secretKey string, token string, role models.UserKind) (*models.User, error) {
	user, err := jwt.DecryptToken(secretKey, token)

	if err != nil {
		return nil, err
	}

	if user.UserKind < role {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  errorHandler.Forbidden.Error(),
		}
	}

	return user, nil
}
