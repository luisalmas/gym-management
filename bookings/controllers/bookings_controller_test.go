package controllers

import (
	"encoding/json"
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/errors"
	"gym-management/bookings/services"
	classesErrors "gym-management/classes/models/errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBookingsController(t *testing.T) {
	var firstbooking = &dtos.BookingCompleteDTO{
		BookingId: 1,
		Name: "Peter",
		Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
		ClassId: 1,
	}
	
	var secondbooking = &dtos.BookingCompleteDTO{
		BookingId: 2,
		Name: "Samantha",
		Date: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
		ClassId: 1,
	}
	
	var bookingToInsert = &dtos.BookingDTO{
		Name: "Dany",
		Date: time.Date(2024, time.January, 27,  0, 0, 0, 0, time.UTC),
		ClassId: 1,
	}
	
	var insertedbooking = &dtos.BookingCompleteDTO{
		BookingId: 4,
		Name: "Dany",
		Date: time.Date(2024, time.January, 27,  0, 0, 0, 0, time.UTC),
		ClassId: 1,
	}

	controller := NewBookingsController()
	router := gin.Default()
	prefix := router.Group("/api")
	
	controller.SetupRoutes(prefix)

	bookingsMock := new(services.MockBookingsService)
	controller.BookingsService = bookingsMock

	//===================== GetBookings tests ==============================================

	t.Run("GetBookings", func(t *testing.T) {
		bookingsMock.On("GetBookings", mock.Anything).Return(&[]dtos.BookingCompleteDTO{*firstbooking,*secondbooking}).Once()

		request, _ := http.NewRequest("GET","/api/bookings",nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		var responseData []dtos.BookingCompleteDTO
		err := json.Unmarshal(response.Body.Bytes(), &responseData)
		assert.Nil(t, err)
		assert.Equal(t, []dtos.BookingCompleteDTO{*firstbooking,*secondbooking}, responseData)
	})

	//===================== GetBooking tests ==============================================

	t.Run("GetBooking", func(t *testing.T) {
		bookingsMock.On("GetBooking", mock.Anything).Return(firstbooking, nil).Once()

		request, _ := http.NewRequest("GET","/api/bookings/1",nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		var responseData dtos.BookingCompleteDTO
		err := json.Unmarshal(response.Body.Bytes(), &responseData)
		assert.Nil(t, err)
		assert.Equal(t, *firstbooking, responseData)
	})

	t.Run("GetBookingInvalidId", func(t *testing.T){

		request, _ := http.NewRequest("GET","/api/bookings/asd", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("GetBookingNotFound", func(t *testing.T){
		bookingsMock.On("GetBooking", mock.Anything).Return(nil, errors.NewBookingNotFoundError()).Once()

		request, _ := http.NewRequest("GET","/api/bookings/1",nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	//===================== PostBookings tests ==============================================

	t.Run("PostBookings", func(t *testing.T){
		bookingsMock.On("InsertNewBooking", mock.Anything).Return(insertedbooking, nil).Once()
		
		jsonData, errorJson := json.Marshal(*bookingToInsert)

		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("POST", "/api/bookings", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var responseData dtos.BookingCompleteDTO
		err := json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Nil(t, err)
		assert.Equal(t, *insertedbooking, responseData)
	})

	t.Run("PostBookingsNoBody", func(t *testing.T){

		request, _ := http.NewRequest("POST", "/api/bookings", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("PostBookingWithError", func(t *testing.T){
		bookingsMock.On("InsertNewBooking", bookingToInsert).Return(nil, errors.NewBookingDateInvalid()).Once()

		jsonData, errorJson := json.Marshal(bookingToInsert)

		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("POST", "/api/bookings", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("PostBookingClassNotFound", func(t *testing.T) {
		bookingsMock.On("InsertNewBooking", bookingToInsert).Return(nil, classesErrors.NewClassNotFoundError()).Once()

		jsonData, errorJson := json.Marshal(bookingToInsert)

		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("POST", "/api/bookings", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	//===================== PutBookings tests ==============================================
	
	t.Run("PutBooking", func(t *testing.T){
		bookingsMock.On("UpdateBooking", mock.AnythingOfType("int"), bookingToInsert).Return(insertedbooking, nil, nil).Once()
		
		jsonData, errorJson := json.Marshal(bookingToInsert)
		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("PUT", "/api/bookings/4", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var responseData dtos.BookingCompleteDTO
		err := json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, *insertedbooking, responseData)
	})

	t.Run("PutBookingNoBody", func(t *testing.T){

		request, _ := http.NewRequest("PUT", "/api/bookings/1", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("PutBookingInvalidId", func(t *testing.T){
		jsonData, errorJson := json.Marshal(bookingToInsert)
		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("PUT", "/api/bookings/asd", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("PutBookingNotFound", func(t *testing.T){
		bookingsMock.On("UpdateBooking", mock.Anything, bookingToInsert).Return(nil, errors.NewBookingNotFoundError()).Once()

		jsonData, errorJson := json.Marshal(*bookingToInsert)
		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("PUT", "/api/bookings/4", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("PutBookingError", func(t *testing.T){
		bookingsMock.On("UpdateBooking", mock.Anything, bookingToInsert).Return(nil, errors.NewBookingDateInvalid()).Once()

		jsonData, errorJson := json.Marshal(bookingToInsert)
		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("PUT", "/api/bookings/4", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)


		require.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("PutBookingClassNotFound", func(t *testing.T) {
		bookingsMock.On("UpdateBooking", mock.Anything, bookingToInsert).Return(nil, classesErrors.NewClassNotFoundError()).Once()

		jsonData, errorJson := json.Marshal(bookingToInsert)
		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("PUT", "/api/bookings/4", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)


		require.Equal(t, http.StatusNotFound, response.Code)
	})

	//===================== DeleteBookings tests ==============================================

	t.Run("DeleteBooking", func(t *testing.T){
		bookingsMock.On("DeleteBooking", mock.Anything).Return(insertedbooking, nil, nil).Once()

		request, _ := http.NewRequest("DELETE", "/api/bookings/1", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var responseData dtos.BookingCompleteDTO
		err := json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, *insertedbooking, responseData)
	})

	t.Run("DeleteBookingInvalidId", func(t *testing.T){

		request, _ := http.NewRequest("DELETE", "/api/bookings/asd", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		require.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("DeleteBookingNotFound", func(t *testing.T){
		bookingsMock.On("DeleteBooking", mock.Anything).Return(nil, errors.NewBookingNotFoundError()).Once()

		request, _ := http.NewRequest("DELETE", "/api/bookings/1", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("DeleteBookingError", func(t *testing.T){
		bookingsMock.On("DeleteBooking", mock.Anything).Return(nil, errors.NewBookingNotFoundError()).Once()

		request, _ := http.NewRequest("DELETE", "/api/bookings/1", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	bookingsMock.AssertExpectations(t)
}