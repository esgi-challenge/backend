package errorHandler

import (
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

var (
	NotFound            = errors.New("Not Found")
	BadRequest          = errors.New("Bad Request")
	BadURLParams        = errors.New("Invalid URL params")
	Unauthorized        = errors.New("You need to login")
	Forbidden           = errors.New("You are not allowed to access this ressource")
	BadBodyParams       = errors.New("Invalid Body params")
	Conflict            = errors.New("Conflict")
	InternalServerError = errors.New("Internal Server Error")
)

type HttpErr interface {
	Status() int
	Error() string
}

type HttpError struct {
	HttpStatus int    `json:"status,omitempty"`
	HttpError  string `json:"error,omitempty"`
}

func (e HttpError) Error() string {
	return fmt.Sprintf(e.HttpError)
}

func (e HttpError) Status() int {
	return e.HttpStatus
}

func NewHttpError(status int, err string) HttpErr {
	return HttpError{
		HttpStatus: status,
		HttpError:  err,
	}
}

func ParseError(err error) HttpErr {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return NewHttpError(http.StatusNotFound, NotFound.Error())
	case errors.Is(err, gorm.ErrInvalidData):
		return NewHttpError(http.StatusBadRequest, NotFound.Error())
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return NewHttpError(http.StatusConflict, Conflict.Error())
	default:
		if httpErr, ok := err.(HttpErr); ok {
			return httpErr
		}
		return NewHttpError(http.StatusInternalServerError, InternalServerError.Error())
	}
}

func ErrorResponse(err error) (int, interface{}) {
	return ParseError(err).Status(), ParseError(err)
}

func UrlParamsErrorResponse() (int, interface{}) {
	err := HttpError{
		HttpStatus: http.StatusBadRequest,
		HttpError:  BadURLParams.Error(),
	}
	return ParseError(err).Status(), ParseError(err)
}

func BodyParamsErrorResponse() (int, interface{}) {
	err := HttpError{
		HttpStatus: http.StatusBadRequest,
		HttpError:  BadBodyParams.Error(),
	}
	return ParseError(err).Status(), ParseError(err)
}

func UnauthorizedErrorResponse() (int, interface{}) {
	err := HttpError{
		HttpStatus: http.StatusUnauthorized,
		HttpError:  Unauthorized.Error(),
	}
	return ParseError(err).Status(), ParseError(err)
}

func ForbiddenErrorResponse() (int, interface{}) {
	err := HttpError{
		HttpStatus: http.StatusForbidden,
		HttpError:  Forbidden.Error(),
	}
	return ParseError(err).Status(), ParseError(err)
}
