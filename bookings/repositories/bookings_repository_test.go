package repositories

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetBookings(t *testing.T) {
	bookingsRepository := NewBookingsRepository()
	bookings := bookingsRepository.GetBookings()

	require.Equal(t, *bookings, []dtos.BookingCompleteDTO {
		{
			Id: 1,
			Name: "Peter",
			Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
		},
		{
			Id: 2,
			Name: "Samantha",
			Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
		},
	})
}

func TestGetBooking(t *testing.T){
	bookingsRepository := NewBookingsRepository()

	validBooking, err1 := bookingsRepository.GetBooking(1)

	require.Nil(t, err1)
	require.Equal(t, *validBooking, entities.Booking{
			Id: 1,
			Name: "Peter",
			Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
		})

	invalidBooking, err2 := bookingsRepository.GetBooking(0)
	require.Nil(t, invalidBooking)
	require.NotNil(t, err2)
}