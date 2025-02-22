package controllers

import (
	"apiGo/src/Empleados/application"
	"apiGo/src/Empleados/domain/entities"

	"github.com/gin-gonic/gin"
)

type AddEmpleadosController struct {
	useCase application.AddEmpleadosUseCase
}

func NewAddEmpleadosController(useCase application.AddEmpleadosUseCase) *AddEmpleadosController {
	return &AddEmpleadosController{useCase: useCase}
}

func (controller *AddEmpleadosController) AddEmpleado(c *gin.Context) {
	var empleado entities.Empleado
	if err := c.BindJSON(&empleado); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.useCase.AddEmpleado(empleado)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}
