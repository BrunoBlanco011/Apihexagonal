package controllers

import (
	"apiGo/src/Empleados/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewAllEmpleadoController struct {
	useCase application.ViewAllEmpleadoUseCase
}

func NewViewAllEmpleadoController(useCase application.ViewAllEmpleadoUseCase) *ViewAllEmpleadoController {
	return &ViewAllEmpleadoController{useCase: useCase}
}

func (controller *ViewAllEmpleadoController) ViewAllEmpleado(c *gin.Context) {
	empleados, err := controller.useCase.GetAllEmpleados()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, empleados)
}
