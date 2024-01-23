package controllers

import (
	"gym-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClassesController struct {
	service *services.ClassesService
}

func NewClassesController() *ClassesController {
	return &ClassesController{
		service: services.NewClassesService(),
	}
}

func (ctrl *ClassesController) SetupRoutes(router *gin.RouterGroup) {
	router.GET("/classes", ctrl.getClassesSchedules)
	//router.POST("/classes", ctrl.service.GetItems)
	//router.PUT("/classes", ctrl.service.GetItems)
}

// GetClasses             godoc
// @Summary      Get classes
// @Description  Returns all scheduled classes.
// @Tags         classes
// @Produce      json
// @Success      200  {array}  entities.ClassSchedule
// @Router       /classes [get]
func (ctrl *ClassesController) getClassesSchedules(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ctrl.service.GetClassesSchedules())
}