package controllers

import "github.com/gin-gonic/gin"

type ClassesControllerInterface interface {
	getClassesSchedules(c *gin.Context)
	getClassSchedule(c *gin.Context)
	postClassSchedule(c *gin.Context)
	putClassSchedule(c *gin.Context)
}