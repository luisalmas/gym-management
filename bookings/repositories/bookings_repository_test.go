package repositories

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookingsRepository(t *testing.T) {
	var firstbooking = &entities.Booking{
			BookingId: 1,
			Name: "Peter",
			Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
			ClassId: 1,
		}

	var secondbooking = &entities.Booking{
			BookingId: 2,
			Name: "Samantha",
			Date: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
			ClassId: 1,
		}

	var bookingToInsert = &entities.Booking{
			BookingId: 3,
			Name: "Jonas",
			Date: time.Date(2024, time.January, 25, 0, 0, 0, 0, time.UTC),
			ClassId: 1,
		}

	type DeleteBookingsFromClassTestParams struct {
			ClassId int
			DateBegin time.Time
			DateEnd time.Time
			ExpectedResult any
		}
	
	deleteBookingsFromClassTests := []DeleteBookingsFromClassTestParams{
			{
				ClassId: 1,
				DateBegin: time.Time{},
				DateEnd: time.Time{},
				ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingCompleteDTO(),*secondbooking.ToBookingCompleteDTO()},
			},
			{
				ClassId: 1,
				DateBegin: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
				DateEnd: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
				ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingCompleteDTO()},
			},
			{
				ClassId: 1,
				DateBegin: time.Now(),
				DateEnd: time.Time{},
				ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingCompleteDTO(), *secondbooking.ToBookingCompleteDTO()},
			},
			{
				ClassId: 1,
				DateBegin: time.Time{},
				DateEnd: time.Now(),
				ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingCompleteDTO(), *secondbooking.ToBookingCompleteDTO()},
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
				ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingCompleteDTO(), *secondbooking.ToBookingCompleteDTO()},
			},
		}

	//===================== GetBookings tests ==============================================

	t.Run("GetBookings", func(t *testing.T) {
		bookingsRepository := NewBookingsRepository()
		bookings := bookingsRepository.GetBookings()

		assert.Equal(t, []dtos.BookingCompleteDTO {
			*firstbooking.ToBookingCompleteDTO(),
			*secondbooking.ToBookingCompleteDTO(),
		}, *bookings)
	})

	//===================== GetBookings tests ==============================================

	t.Run("GetBooking", func(t *testing.T) {
		bookingsRepository := NewBookingsRepository()
		booking, err := bookingsRepository.GetBooking(1)

		assert.Nil(t, err)
		assert.Equal(t, firstbooking, booking)
	})

	t.Run("GetBookingNotFound", func(t *testing.T) {
		bookingsRepository := NewBookingsRepository()
		booking, err := bookingsRepository.GetBooking(0)

		assert.NotNil(t, err)
		assert.Nil(t, booking)
	})

	//===================== GetBookingsFromClass tests ==============================================

	t.Run("GetBookingsFromClass", func(t *testing.T){
		bookingsRepository := NewBookingsRepository()
		bookings := bookingsRepository.GetBookingsFromClass(1, time.Time{})

		assert.Equal(t, []dtos.BookingCompleteDTO{*firstbooking.ToBookingCompleteDTO(), *secondbooking.ToBookingCompleteDTO()}, *bookings)
	})

	t.Run("GetBookingsFromClassWithDate", func(t *testing.T){
		bookingsRepository := NewBookingsRepository()
		bookings := bookingsRepository.GetBookingsFromClass(1, time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC))

		assert.Equal(t, []dtos.BookingCompleteDTO{*secondbooking.ToBookingCompleteDTO()}, *bookings)
	})

	t.Run("GetBookingsFromClassNotFound", func(t *testing.T){
		bookingsRepository := NewBookingsRepository()
		bookings := bookingsRepository.GetBookingsFromClass(0, time.Time{})

		assert.Equal(t, []dtos.BookingCompleteDTO{}, *bookings)
	})

	t.Run("GetBookingsFromDateNotFound", func(t *testing.T){
		bookingsRepository := NewBookingsRepository()
		bookings := bookingsRepository.GetBookingsFromClass(1, time.Date(2024, time.December, 26,  0, 0, 0, 0, time.UTC))

		assert.Equal(t, []dtos.BookingCompleteDTO{}, *bookings)
	})

	//===================== InsertNewBooking tests ==============================================

	t.Run("InsertNewBooking", func(t *testing.T){
		bookingsRepository := NewBookingsRepository()
		insertedBooking := bookingsRepository.InsertNewBooking(bookingToInsert)

		assert.Equal(t, bookingToInsert.ToBookingCompleteDTO(), insertedBooking)

		booking, err := bookingsRepository.GetBooking(3)

		assert.Nil(t, err)
		assert.Equal(t, bookingToInsert, booking)
	})

	//===================== UpdateBooking tests ==============================================

	t.Run("UpdateBooking", func(t *testing.T){
		bookingsRepository := NewBookingsRepository()
		updatedBooking := bookingsRepository.UpdateBooking(1, bookingToInsert)

		assert.Equal(t, dtos.BookingCompleteDTO{
			BookingId: 1,
			Name: "Jonas",
			Date: time.Date(2024, time.January, 25, 0, 0, 0, 0, time.UTC),
			ClassId: 1,
		}, *updatedBooking)
	})

	//===================== DeleteBooking tests ==============================================

	t.Run("DeleteBooking", func(t *testing.T){
		bookingsRepository := NewBookingsRepository()
		deletedBooking, err := bookingsRepository.DeleteBooking(1)

		assert.Nil(t, err)
		assert.Equal(t, firstbooking.ToBookingCompleteDTO(), deletedBooking)

		booking, err := bookingsRepository.GetBooking(1)

		assert.NotNil(t, err)
		assert.Nil(t, booking)
	})

	t.Run("DeleteBookingNotFound", func(t *testing.T){
		bookingsRepository := NewBookingsRepository()
		deletedBooking, err := bookingsRepository.DeleteBooking(0)

		assert.NotNil(t, err)
		assert.Nil(t, deletedBooking)
	})

	//===================== DeleteBookingsFromClass tests ==============================================

	for index, test := range deleteBookingsFromClassTests {
		t.Run("DeleteBookingsFromClass_" + strconv.Itoa(index), func(t *testing.T) {
			bookingsRepository := NewBookingsRepository()
			result := bookingsRepository.DeleteBookingsFromClass(test.ClassId, test.DateBegin, test.DateEnd)
			require.Equal(t, test.ExpectedResult, *result)
		})
	}

}

// func TestGetBookings(t *testing.T) {
// 	bookingsRepository := NewBookingsRepository()
// 	bookings := bookingsRepository.GetBookings()

// 	assert.Equal(t, []dtos.BookingCompleteDTO {
// 		*firstbooking.ToBookingCompleteDTO(),
// 		*secondbooking.ToBookingCompleteDTO(),
// 	}, *bookings)
// }

// func TestGetBooking(t *testing.T){
// 	bookingsRepository := NewBookingsRepository()

// 	validBooking, err1 := bookingsRepository.GetBooking(1)

// 	require.Nil(t, err1)
// 	require.Equal(t, firstbooking, *validBooking)

// 	invalidBooking, err2 := bookingsRepository.GetBooking(0)
// 	require.Nil(t, invalidBooking)
// 	require.NotNil(t, err2)
// }

// func TestGetBookingFromClass(t *testing.T){
// 	bookingsRepository := NewBookingsRepository()

// 	bookingsFromClass := bookingsRepository.GetBookingsFromClass(1 ,time.Time{})
// 	require.Equal(t, *bookingsFromClass, []dtos.BookingCompleteDTO {
// 		*firstbooking.ToBookingCompleteDTO(),
// 		*secondbooking.ToBookingCompleteDTO(),
// 	})

// 	bookingsFromClassWithDate := bookingsRepository.GetBookingsFromClass(1 ,time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC))
// 	require.Equal(t, []dtos.BookingCompleteDTO {
// 		*firstbooking.ToBookingCompleteDTO(),
// 	}, *bookingsFromClassWithDate)

// 	testNoClass := bookingsRepository.GetBookingsFromClass(0, time.Time{})
// 	require.Equal(t, []dtos.BookingCompleteDTO{}, *testNoClass)

// 	testNoClassWithDate := bookingsRepository.GetBookingsFromClass(0, time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC))
// 	require.Equal(t, []dtos.BookingCompleteDTO{}, *testNoClassWithDate)

// 	testDateOusideClassRange := bookingsRepository.GetBookingsFromClass(1, time.Date(2024, time.February, 25,  0, 0, 0, 0, time.UTC))
// 	require.Equal(t, []dtos.BookingCompleteDTO{}, *testDateOusideClassRange)
// }

// func TestInsertNewBooking(t *testing.T) {
// 	bookingsRepository := NewBookingsRepository()

// 	bookingToInsert := entities.Booking{
// 		BookingId: 0,
// 		Name: "Samuel",
// 		Date: time.Date(2024, time.January, 27,  0, 0, 0, 0, time.UTC),
// 		ClassId: 2,
// 	}

// 	insertedBooking := bookingsRepository.InsertNewBooking(&bookingToInsert)

// 	require.Equal(t, bookingToInsert.ToBookingCompleteDTO(), insertedBooking)
// }

// func TestUpdateBooking(t *testing.T) {
// 	bookingsRepository := NewBookingsRepository()

// 	bookingToUpdate := entities.Booking{
// 		BookingId: 1,
// 		Name: "Sam",
// 		Date: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
// 		ClassId: 1,
// 	}

// 	updatedBooking := bookingsRepository.UpdateBooking(bookingToUpdate.BookingId, &bookingToUpdate)

// 	require.Equal(t, bookingToUpdate.ToBookingCompleteDTO(), updatedBooking)
// }

// func TestDeleteBooking(t *testing.T) {
// 	bookingsRepository := NewBookingsRepository()

// 	deletedBooking, err := bookingsRepository.DeleteBooking(firstbooking.BookingId)
// 	require.Nil(t, err)
// 	require.Equal(t, firstbooking.ToBookingCompleteDTO(), deletedBooking)

// 	deleteBookingFail, errFail := bookingsRepository.DeleteBooking(0)
// 	require.NotNil(t, errFail)
// 	require.Nil(t, deleteBookingFail)
// 	//TODO perform get to confirm delete?
// }

// type DeleteBookingsFromClassTestParams struct {
// 	ClassId int
// 	DateBegin time.Time
// 	DateEnd time.Time
// 	ExpectedResult any
// }

// func TestDeleteBookingsFromClass(t *testing.T){
// 	tests := []DeleteBookingsFromClassTestParams{
// 		{
// 			ClassId: 1,
// 			DateBegin: time.Time{},
// 			DateEnd: time.Time{},
// 			ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingCompleteDTO(),*secondbooking.ToBookingCompleteDTO()},
// 		},
// 		{
// 			ClassId: 1,
// 			DateBegin: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
// 			DateEnd: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
// 			ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingCompleteDTO()},
// 		},
// 		{
// 			ClassId: 1,
// 			DateBegin: time.Now(),
// 			DateEnd: time.Time{},
// 			ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingCompleteDTO(), *secondbooking.ToBookingCompleteDTO()},
// 		},
// 		{
// 			ClassId: 1,
// 			DateBegin: time.Time{},
// 			DateEnd: time.Now(),
// 			ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingCompleteDTO(), *secondbooking.ToBookingCompleteDTO()},
// 		},
// 		{
// 			ClassId: 0,
// 			DateBegin: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
// 			DateEnd: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
// 			ExpectedResult: []dtos.BookingCompleteDTO {},
// 		},
// 		{
// 			ClassId: 1,
// 			DateBegin: time.Date(2024, time.February, 01,  0, 0, 0, 0, time.UTC),
// 			DateEnd: time.Date(2024, time.February, 26,  0, 0, 0, 0, time.UTC),
// 			ExpectedResult: []dtos.BookingCompleteDTO {*firstbooking.ToBookingCompleteDTO(), *secondbooking.ToBookingCompleteDTO()},
// 		},
// 	}

// 	for _, test := range tests {
// 		bookingsRepository := NewBookingsRepository()
// 		result := bookingsRepository.DeleteBookingsFromClass(test.ClassId, test.DateBegin, test.DateEnd)
// 		require.Equal(t, test.ExpectedResult, *result)
// 	}
// }