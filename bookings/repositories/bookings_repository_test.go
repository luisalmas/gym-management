package repositories

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var firstbooking = entities.Booking{
		BookingId: 1,
		Name: "Peter",
		Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
		ClassId: 1,
	}

var secondbooking = entities.Booking{
		BookingId: 2,
		Name: "Samantha",
		Date: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
		ClassId: 1,
	}

func TestGetBookings(t *testing.T) {
	bookingsRepository := NewBookingsRepository()
	bookings := bookingsRepository.GetBookings()

	require.Equal(t, []dtos.BookingCompleteDTO {
		*firstbooking.ToBookingDTO(),
		*secondbooking.ToBookingDTO(),
	}, *bookings)
}

func TestGetBooking(t *testing.T){
	bookingsRepository := NewBookingsRepository()

	validBooking, err1 := bookingsRepository.GetBooking(1)

	require.Nil(t, err1)
	require.Equal(t, firstbooking, *validBooking)

	invalidBooking, err2 := bookingsRepository.GetBooking(0)
	require.Nil(t, invalidBooking)
	require.NotNil(t, err2)
}

func TestGetBookingFromClass(t *testing.T){
	bookingsRepository := NewBookingsRepository()

	bookingsFromClass := bookingsRepository.GetBookingsFromClass(1 ,time.Time{})
	require.Equal(t, *bookingsFromClass, []dtos.BookingCompleteDTO {
		*firstbooking.ToBookingDTO(),
		*secondbooking.ToBookingDTO(),
	})

	bookingsFromClassWithDate := bookingsRepository.GetBookingsFromClass(1 ,time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC))
	require.Equal(t, []dtos.BookingCompleteDTO {
		*firstbooking.ToBookingDTO(),
	}, *bookingsFromClassWithDate)

	testNoClass := bookingsRepository.GetBookingsFromClass(0, time.Time{})
	require.Equal(t, []dtos.BookingCompleteDTO{}, *testNoClass)

	testNoClassWithDate := bookingsRepository.GetBookingsFromClass(0, time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC))
	require.Equal(t, []dtos.BookingCompleteDTO{}, *testNoClassWithDate)

	testDateOusideClassRange := bookingsRepository.GetBookingsFromClass(1, time.Date(2024, time.February, 25,  0, 0, 0, 0, time.UTC))
	require.Equal(t, []dtos.BookingCompleteDTO{}, *testDateOusideClassRange)
}

func TestInsertNewBooking(t *testing.T) {
	bookingsRepository := NewBookingsRepository()

	bookingToInsert := entities.Booking{
		BookingId: 0,
		Name: "Samuel",
		Date: time.Date(2024, time.January, 27,  0, 0, 0, 0, time.UTC),
		ClassId: 2,
	}

	insertedBooking := bookingsRepository.InsertNewBooking(&bookingToInsert)

	require.Equal(t, bookingToInsert.ToBookingDTO(), insertedBooking)
}

func TestUpdateBooking(t *testing.T) {
	bookingsRepository := NewBookingsRepository()

	bookingToUpdate := entities.Booking{
		BookingId: 1,
		Name: "Sam",
		Date: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
		ClassId: 1,
	}

	updatedBooking := bookingsRepository.UpdateBooking(bookingToUpdate.BookingId, &bookingToUpdate)

	require.Equal(t, bookingToUpdate.ToBookingDTO(), updatedBooking)
}

func TestDeleteBooking(t *testing.T) {
	bookingsRepository := NewBookingsRepository()

	deletedBooking, err := bookingsRepository.DeleteBooking(firstbooking.BookingId)
	require.Nil(t, err)
	require.Equal(t, firstbooking.ToBookingDTO(), deletedBooking)

	deleteBookingFail, errFail := bookingsRepository.DeleteBooking(0)
	require.NotNil(t, errFail)
	require.Nil(t, deleteBookingFail)
	//TODO perform get to confirm delete?
}

type DeleteBookingsFromClassTestParams struct {
	ClassId int
	DateBegin time.Time
	DateEnd time.Time
	ExpectedResult any
}

func TestDeleteBookingsFromClass(t *testing.T){
	tests := []DeleteBookingsFromClassTestParams{
		{
			ClassId: 1,
			DateBegin: time.Time{},
			DateEnd: time.Time{},
			ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingDTO(),*secondbooking.ToBookingDTO()},
		},
		{
			ClassId: 1,
			DateBegin: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
			DateEnd: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
			ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingDTO()},
		},
		{
			ClassId: 1,
			DateBegin: time.Now(),
			DateEnd: time.Time{},
			ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingDTO(), *secondbooking.ToBookingDTO()},
		},
		{
			ClassId: 1,
			DateBegin: time.Time{},
			DateEnd: time.Now(),
			ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingDTO(), *secondbooking.ToBookingDTO()},
		},
		{
			ClassId: 0,
			DateBegin: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
			DateEnd: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
			ExpectedResult: []dtos.BookingCompleteDTO {},
		},
		{
			ClassId: 1,
			DateBegin: time.Date(2024, time.February, 01,  0, 0, 0, 0, time.UTC),
			DateEnd: time.Date(2024, time.February, 26,  0, 0, 0, 0, time.UTC),
			ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingDTO(), *secondbooking.ToBookingDTO()},
		},
	}

	for _, test := range tests {
		bookingsRepository := NewBookingsRepository()
		result := bookingsRepository.DeleteBookingsFromClass(test.ClassId, test.DateBegin, test.DateEnd)
		require.Equal(t, test.ExpectedResult, *result)
	}
}