package myerrors

import (
	"errors"
	"fmt"
)

type MongoError struct {
	Message string
	Code    int
}

type MyError struct {
	Code    int
	Message string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("%s with code %d", err.Message, err.Code)
}

var ErrDocumentNotFound = errors.New("document not found")

var ErrInternalServerError = errors.New("internel server error")

func (m *MongoError) Error() string {
	return m.Message
}

type ApiError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(status int, message string) *ApiError {

	return &ApiError{
		Code:    status,
		Message: message,
	}
}

func NotFoundError(message string) *ApiError {
	return NewApiError(404, message)
}

func InternalServerError(message string) *ApiError {
	return NewApiError(500, message)
}
func RecordExistsError(record string) *ApiError {
	return NewApiError(400, record+" already exists.")
}
