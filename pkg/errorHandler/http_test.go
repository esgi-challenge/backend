package errorHandler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewHttpError(t *testing.T) {
	err := NewHttpError(http.StatusNotFound, NotFound.Error())
	assert.Equal(t, http.StatusNotFound, err.Status())
	assert.Equal(t, "Not Found", err.Error())
}

func TestParseError(t *testing.T) {
	tests := []struct {
		name          string
		inputError    error
		expectedError HttpErr
	}{
		{
			name:          "Record Not Found",
			inputError:    gorm.ErrRecordNotFound,
			expectedError: NewHttpError(http.StatusNotFound, NotFound.Error()),
		},
		{
			name:          "Invalid Data",
			inputError:    gorm.ErrInvalidData,
			expectedError: NewHttpError(http.StatusBadRequest, NotFound.Error()),
		},
		{
			name:          "Duplicated Key",
			inputError:    gorm.ErrDuplicatedKey,
			expectedError: NewHttpError(http.StatusConflict, Conflict.Error()),
		},
		{
			name:          "HttpErr",
			inputError:    NewHttpError(http.StatusUnauthorized, Unauthorized.Error()),
			expectedError: NewHttpError(http.StatusUnauthorized, Unauthorized.Error()),
		},
		{
			name:          "Internal Server Error",
			inputError:    errors.New("unknown error"),
			expectedError: NewHttpError(http.StatusInternalServerError, InternalServerError.Error()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ParseError(tt.inputError)
			assert.Equal(t, tt.expectedError.Status(), err.Status())
			assert.Equal(t, tt.expectedError.Error(), err.Error())
		})
	}
}

func TestErrorResponse(t *testing.T) {
	err := gorm.ErrRecordNotFound
	status, response := ErrorResponse(err)
	expectedError := NewHttpError(http.StatusNotFound, NotFound.Error())

	assert.Equal(t, expectedError.Status(), status)
	assert.Equal(t, expectedError.Error(), response.(HttpError).Error())
}

func TestUrlParamsErrorResponse(t *testing.T) {
	status, response := UrlParamsErrorResponse()
	expectedError := NewHttpError(http.StatusBadRequest, BadURLParams.Error())

	assert.Equal(t, expectedError.Status(), status)
	assert.Equal(t, expectedError.Error(), response.(HttpError).Error())
}

func TestBodyParamsErrorResponse(t *testing.T) {
	status, response := BodyParamsErrorResponse()
	expectedError := NewHttpError(http.StatusBadRequest, BadBodyParams.Error())

	assert.Equal(t, expectedError.Status(), status)
	assert.Equal(t, expectedError.Error(), response.(HttpError).Error())
}

func TestUnauthorizedErrorResponse(t *testing.T) {
	status, response := UnauthorizedErrorResponse()
	expectedError := NewHttpError(http.StatusUnauthorized, Unauthorized.Error())

	assert.Equal(t, expectedError.Status(), status)
	assert.Equal(t, expectedError.Error(), response.(HttpError).Error())
}

func TestForbiddenErrorResponse(t *testing.T) {
	status, response := ForbiddenErrorResponse()
	expectedError := NewHttpError(http.StatusForbidden, Forbidden.Error())

	assert.Equal(t, expectedError.Status(), status)
	assert.Equal(t, expectedError.Error(), response.(HttpError).Error())
}

func TestInternalServerErrorResponse(t *testing.T) {
	status, response := InternalServerErrorResponse()
	expectedError := NewHttpError(http.StatusInternalServerError, InternalServerError.Error())

	assert.Equal(t, expectedError.Status(), status)
	assert.Equal(t, expectedError.Error(), response.(HttpError).Error())
}
