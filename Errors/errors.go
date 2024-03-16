package myerrors

import "errors"

type MongoError struct {
	Message string
	Code    int
}

var ErrDocumentNotFound = errors.New("document not found")

var ErrInternalServerError = errors.New("internel server error")

func (m *MongoError) Error() string {
	return m.Message
}

type ApiError struct {
	Message string `json:"message"`
	Code    int
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
