package controllers

import (
	"apiGo/src/Empleados/application"
	"apiGo/src/Empleados/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEmpleadoController struct {
	useCase application.UpdateEmpleadoUseCase
}

func NewUpdateEmpleadoController(useCase application.UpdateEmpleadoUseCase) *UpdateEmpleadoController {
	return &UpdateEmpleadoController{useCase: useCase}
}

func (controller *UpdateEmpleadoController) UpdateEmpleado(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var empleado entities.Empleado
	if err := c.BindJSON(&empleado); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	empleado.Id = id
	result, err := controller.useCase.UpdateEmpleado(empleado)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}
