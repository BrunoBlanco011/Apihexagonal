package controllers

import (
	"apiGo/src/Empleados/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteEmpleadoController struct {
	useCase *application.DeleteEmpleadoUseCase
}

func NewDeleteEmpleadoController(useCase application.DeleteEmpleadoUseCase) *DeleteEmpleadoController {
	return &DeleteEmpleadoController{useCase: &useCase}
}

func (controller *DeleteEmpleadoController) DeleteEmpleado(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = controller.useCase.DeleteEmpleado(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Empleado eliminado"})
}
