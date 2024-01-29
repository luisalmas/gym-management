package controllers

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingsController struct {
	BookingsService services.BookingsServiceInterface
}

func NewBookingsController() *BookingsController {
	return &BookingsController{
		BookingsService: services.NewBookingsService(),
	}
}

func (ctrl *BookingsController) SetupRoutes(router *gin.RouterGroup){
	router.GET("/bookings", ctrl.getBookings)
	router.GET("/bookings/:id", ctrl.getBooking)
	router.POST("/bookings", ctrl.postBooking)
	router.PUT("/bookings/:id", ctrl.putBooking)
	router.DELETE("/bookings/:id", ctrl.deleteBooking)
}

// GetCBookings             godoc
// @Summary      Get bookings
// @Description  Returns all bookings.
// @Tags         bookings
// @Produce      json
// @Success      200  {array}  dtos.BookingCompleteDTO
// @Router       /bookings [get]
func (ctrl *BookingsController) getBookings(c *gin.Context){
	c.IndentedJSON(http.StatusOK, ctrl.BookingsService.GetBookings())
}

// GetBooking            godoc
// @Summary      Get booking
// @Description  Returns single booking.
// @Tags         bookings
// @Produce      json
//@Param         id  path      string true  "Booking Id"
// @Success      200  {object}  dtos.BookingCompleteDTO
// @Failure      404
// @Router       /bookings/{id} [get]
func (ctrl *BookingsController) getBooking(c *gin.Context){
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, idError.Error())
		return
	}

	booking, err := ctrl.BookingsService.GetBooking(id)
	
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, booking)
}

// PostBookings             godoc
// @Summary      Post booking
// @Description  Post a new booking.
// @Tags         bookings
// @Produce      json
//@Param         bookingDTO  body      dtos.BookingDTO true  "BookingDTO JSON"
// @Success      201 {object} dtos.BookingCompleteDTO
// @Failure      400 
// @Router       /bookings [post]
func (ctrl *BookingsController) postBooking(c *gin.Context) {
	var booking dtos.BookingDTO

	if err := c.BindJSON(&booking); err != nil {
        c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
    }

	insertedBooking, err := ctrl.BookingsService.InsertNewBooking(&booking)

	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, insertedBooking)
}

// PutBookings             godoc
// @Summary      Put booking
// @Description  Updates a booking.
// @Tags         bookings
// @Produce      json
//@Param         id  path      string true  "Booking Id"
//@Param         bookingDTO  body      dtos.BookingDTO true  "BookingDTO JSON"
// @Success      200 {object} dtos.BookingCompleteDTO
// @Failure      400 
// @Failure      404 
// @Router       /bookings/{id} [put]
func (ctrl *BookingsController) putBooking(c *gin.Context) {
	var booking dtos.BookingDTO

	if err := c.BindJSON(&booking); err != nil {
        c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
    }

	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, idError.Error())
		return
	}

	updatedBooking, unfoundError, updateError := ctrl.BookingsService.UpdateBooking(id, &booking)

	if unfoundError != nil{
		c.IndentedJSON(http.StatusNotFound, unfoundError.Error())
		return
	}

	if updateError != nil{
		c.IndentedJSON(http.StatusBadRequest, updateError.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, updatedBooking)
}

// DeleteBookings             godoc
// @Summary      Delete booking
// @Description  Deletes a booking.
// @Tags         bookings
// @Produce      json
//@Param         id  path      string true  "Booking Id"
// @Success      200 {object} dtos.BookingCompleteDTO
// @Failure      400 
// @Failure      404 
// @Router       /bookings/{id} [delete]
func (ctrl *BookingsController) deleteBooking(c *gin.Context) {
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, idError.Error())
		return
	}

	deletedBooking, unfoundError, removeError := ctrl.BookingsService.DeleteBooking(id)

	if unfoundError != nil{
		c.IndentedJSON(http.StatusNotFound, unfoundError.Error())
		return
	}

	if removeError != nil{
		c.IndentedJSON(http.StatusBadRequest, removeError.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, deletedBooking)
}