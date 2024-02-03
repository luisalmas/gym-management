package errors

import "net/http"

type BookingsError struct {
	Message string
	Code    int
}

func (err *BookingsError) Error() string {
	return err.Message
}

func (err *BookingsError) ErrorCode() int {
	return err.Code
}

func NewBookingNotFoundError() *BookingsError {
	return &BookingsError{
		Message: "booking not found",
		Code:    http.StatusNotFound,
	}
}

func NewBookingDateInvalid() *BookingsError{
	return &BookingsError{
		Message: "booking date does not correspond to the specified class",
		Code:    http.StatusBadRequest,
	}
}