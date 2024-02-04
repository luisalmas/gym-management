package controllers

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/services"
	"gym-management/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
)

var (
	g = galidator.New()
	customizer = g.Validator(dtos.BookingDTO{})
  )

type BookingsControllerImpl struct {
	BookingsService services.BookingsService
}

func NewBookingsController() *BookingsControllerImpl {
	return &BookingsControllerImpl{
		BookingsService: services.NewBookingsService(),
	}
}

func (ctrl *BookingsControllerImpl) SetupRoutes(router *gin.RouterGroup){
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
func (ctrl *BookingsControllerImpl) getBookings(c *gin.Context){
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
func (ctrl *BookingsControllerImpl) getBooking(c *gin.Context){
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": idError.Error()})
		return
	}

	booking, err := ctrl.BookingsService.GetBooking(id)
	
	if err != nil {
		utils.ErrorHandler(c, err)
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
func (ctrl *BookingsControllerImpl) postBooking(c *gin.Context) {
	var booking dtos.BookingDTO

	if err := c.BindJSON(&booking); err != nil {

		if e, _ := err.(*time.ParseError); e != nil {
		  	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid date format"})
		  	return
		}

        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": customizer.DecryptErrors(err)})
		return
    }

	insertedBooking, err := ctrl.BookingsService.InsertNewBooking(&booking)

	if err != nil{
		utils.ErrorHandler(c, err)
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
func (ctrl *BookingsControllerImpl) putBooking(c *gin.Context) {
	var booking dtos.BookingDTO

	if err := c.BindJSON(&booking); err != nil {

		if e, _ := err.(*time.ParseError); e != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid date format"})
			return
		}

        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": customizer.DecryptErrors(err)})
		return
    }

	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, idError.Error())
		return
	}

	updatedBooking, err := ctrl.BookingsService.UpdateBooking(id, &booking)

	if err != nil {
		utils.ErrorHandler(c, err)
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
func (ctrl *BookingsControllerImpl) deleteBooking(c *gin.Context) {
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": idError.Error()})
		return
	}

	deletedBooking, err := ctrl.BookingsService.DeleteBooking(id)

	if err != nil{
		utils.ErrorHandler(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, deletedBooking)
}