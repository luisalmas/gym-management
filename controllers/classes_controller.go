package controllers

import (
	"gym-management/services"

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

func (ctrl *ClassesController) SetupRoutes(router *gin.Engine) {
	router.GET("/classes", ctrl.service.GetItems)
	//router.POST("/classes", ctrl.service.GetItems)
	//router.PUT("/classes", ctrl.service.GetItems)
}