package controllers

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockBookingsService struct {
	mock.Mock
}

func (mock *MockBookingsService) GetBookings() error {
	args := mock.Called()
	return args.Error(0)
}

func TestGetBookings(t *testing.T) {
	mockBookingsService := new(MockBookingsService)
	mockBookingsService.On("GetBookings").Return(nil)
}