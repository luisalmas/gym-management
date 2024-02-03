package errors

import "net/http"

type ClassesError struct {
	Message string
	Code    int
}

func (err *ClassesError) Error() string {
	return err.Message
}

func (err *ClassesError) ErrorCode() int {
	return err.Code
}

func NewClassNotFoundError() *ClassesError {
	return &ClassesError{
		Message: "class not found",
		Code:    http.StatusNotFound,
	}
}

func NewClassInvalidError(msg string) *ClassesError {
	return &ClassesError{
		Message: msg,
		Code:    http.StatusBadRequest,
	}
}