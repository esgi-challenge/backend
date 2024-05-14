package errorHandler

import (
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

var (
	NotFound            = errors.New("Not Found")
	BadURLParams        = errors.New("Invalid URL params")
	BadBodyParams       = errors.New("Invalid Body params")
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
