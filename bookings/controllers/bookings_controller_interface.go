package controllers

import "github.com/gin-gonic/gin"

type BookingsControllerInterface interface {
	getBookings(c *gin.Context)
	getBooking(c *gin.Context)
	postBooking(c *gin.Context)
	putBooking(c *gin.Context)
	deleteBooking(c *gin.Context)
}