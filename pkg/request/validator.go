package request

import (
	"errors"
	"strings"

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

