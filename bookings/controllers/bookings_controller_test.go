package controllers

import (
	"encoding/json"
	"errors"
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockBookingsService struct {
	mock.Mock
}

func (mock *MockBookingsService) GetBookings() error {
	args := mock.Called()
	return args.Error(0)
}

var firstbooking = dtos.BookingCompleteDTO{
	BookingId: 1,
	Name: "Peter",
	Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
	ClassId: 1,
}

var secondbooking = dtos.BookingCompleteDTO{
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

func setupController() (*BookingsControllerImpl, *gin.Engine) {
	controller := NewBookingsController()
	router := gin.Default()
	prefix := router.Group("/api")
	
	controller.SetupRoutes(prefix)

	return controller, router
}

func TestGetBookings(t *testing.T) {
	controller, router := setupController()

	bookingsMock := new(services.MockBookingsService)
	bookingsMock.On("GetBookings").Return(&[]dtos.BookingCompleteDTO{
		firstbooking,
		secondbooking,
	})
	controller.BookingsService = bookingsMock

	req, _ := http.NewRequest("GET", "/api/bookings", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
	var response []dtos.BookingCompleteDTO
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Equal(t, []dtos.BookingCompleteDTO{
		firstbooking,
		secondbooking,
	}, response)
}

func TestGetBooking(t *testing.T) {
	controller, router := setupController()

	bookingsMock := new(services.MockBookingsService)
	bookingsMock.On("GetBooking", mock.AnythingOfType("int")).Return(&firstbooking, nil)
	controller.BookingsService = bookingsMock

	req, _ := http.NewRequest("GET", "/api/bookings/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
	var response dtos.BookingCompleteDTO
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	require.Nil(t, err)
	require.Equal(t, firstbooking, response)

	bookingsMock = new(services.MockBookingsService)
	bookingsMock.On("GetBooking", mock.AnythingOfType("int")).Return(nil, errors.New("error"))
	controller.BookingsService = bookingsMock

	req, _ = http.NewRequest("GET", "/api/bookings/99", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusNotFound, resp.Code)

	req, _ = http.NewRequest("GET", "/api/bookings/asd", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestPostBookings(t *testing.T) {
	controller, router := setupController()

	req, _ := http.NewRequest("POST", "/api/bookings", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusBadRequest, resp.Code)

	bookingsMock := new(services.MockBookingsService)
	bookingsMock.On("InsertNewBooking", bookingToInsert).Return(insertedbooking, nil)
	controller.BookingsService = bookingsMock

	jsonData, errorJson := json.Marshal(bookingToInsert)
	require.Nil(t, errorJson)
	req, _ = http.NewRequest("POST", "/api/bookings", strings.NewReader(string(jsonData)))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusCreated, resp.Code)
	var response dtos.BookingCompleteDTO
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	require.Nil(t, err)
	require.Equal(t, *insertedbooking, response)

	bookingsMock = new(services.MockBookingsService)
	bookingsMock.On("InsertNewBooking", bookingToInsert).Return(nil, errors.New("error"))
	controller.BookingsService = bookingsMock

	jsonData, errorJson = json.Marshal(bookingToInsert)
	require.Nil(t, errorJson)
	req, _ = http.NewRequest("POST", "/api/bookings", strings.NewReader(string(jsonData)))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestPutBookings(t *testing.T){
	controller, router := setupController()

	req, _ := http.NewRequest("PUT", "/api/bookings/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusBadRequest, resp.Code)

	jsonData, errorJson := json.Marshal(bookingToInsert)
	require.Nil(t, errorJson)
	req, _ = http.NewRequest("PUT", "/api/bookings/sdf", strings.NewReader(string(jsonData)))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusBadRequest, resp.Code)

	bookingsMock := new(services.MockBookingsService)
	bookingsMock.On("UpdateBooking", mock.AnythingOfType("int"), bookingToInsert).Return(insertedbooking, nil, nil)
	controller.BookingsService = bookingsMock

	jsonData, errorJson = json.Marshal(bookingToInsert)
	require.Nil(t, errorJson)
	req, _ = http.NewRequest("PUT", "/api/bookings/4", strings.NewReader(string(jsonData)))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	var response dtos.BookingCompleteDTO
	err := json.Unmarshal(resp.Body.Bytes(), &response)

	require.Nil(t, err)
	require.Equal(t, http.StatusOK, resp.Code)
	require.Equal(t, *insertedbooking, response)

	bookingsMock = new(services.MockBookingsService)
	bookingsMock.On("UpdateBooking", mock.AnythingOfType("int"), bookingToInsert).Return(nil, errors.New("error"), nil)
	controller.BookingsService = bookingsMock

	jsonData, errorJson = json.Marshal(bookingToInsert)
	require.Nil(t, errorJson)
	req, _ = http.NewRequest("PUT", "/api/bookings/4", strings.NewReader(string(jsonData)))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusNotFound, resp.Code)

	bookingsMock = new(services.MockBookingsService)
	bookingsMock.On("UpdateBooking", mock.AnythingOfType("int"), bookingToInsert).Return(nil, nil, errors.New("error"))
	controller.BookingsService = bookingsMock

	jsonData, errorJson = json.Marshal(bookingToInsert)
	require.Nil(t, errorJson)
	req, _ = http.NewRequest("PUT", "/api/bookings/4", strings.NewReader(string(jsonData)))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestDeleteBookings(t *testing.T){
	controller, router := setupController()

	response := testHttpRequest("DELETE", "/api/bookings/asd", nil, router)

	require.Equal(t, http.StatusBadRequest, response.Code)

	setBookingsMock(
		controller,
		"DeleteBooking",
		[]any{mock.AnythingOfType("int")},
		[]any{insertedbooking, nil, nil},
	)

	response = testHttpRequest("DELETE", "/api/bookings/1", nil, router)

	var responseData dtos.BookingCompleteDTO
	err := json.Unmarshal(response.Body.Bytes(), &responseData)

	require.Nil(t, err)
	require.Equal(t, http.StatusOK, response.Code)
	require.Equal(t, *insertedbooking, responseData)

	setBookingsMock(
		controller,
		"DeleteBooking",
		[]any{mock.AnythingOfType("int")},
		[]any{nil, errors.New("error"), nil},
	)

	response = testHttpRequest("DELETE", "/api/bookings/1", nil, router)

	require.Equal(t, http.StatusNotFound, response.Code)

	setBookingsMock(
		controller,
		"DeleteBooking",
		[]any{mock.AnythingOfType("int")},
		[]any{nil, nil, errors.New("error")},
	)

	response = testHttpRequest("DELETE", "/api/bookings/1", nil, router)

	require.Equal(t, http.StatusBadRequest, response.Code)
}

func testHttpRequest(verb string, uri string, body any, router *gin.Engine) (*httptest.ResponseRecorder) {
	request, _ := http.NewRequest(verb, uri, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	return response
}

func setBookingsMock(controller *BookingsControllerImpl, functionName string, args []any, returns []any){
	bookingsMock := new(services.MockBookingsService)
	bookingsMock.On(functionName, args...).Return(returns...)
	controller.BookingsService = bookingsMock
}