package controllers

import (
	"apiGo/src/Empleados/application"
	
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ViewEmpleadoIDController struct {
	useCase application.ViewEmpleadoIDUseCase
}

func NewViewEmpleadoIDController(useCase application.ViewEmpleadoIDUseCase) *ViewEmpleadoIDController {
	return &ViewEmpleadoIDController{useCase: useCase}
}

func (controller *ViewEmpleadoIDController) ViewEmpleadoID(c *gin.Context) {
    // Obtiene el parámetro 'id' de la URL y lo convierte a entero
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    // Llama al caso de uso para obtener el empleado por ID
    empleado, err := controller.useCase.GetEmpleado(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Empleado no encontrado"})
        return
    }

    // Retorna el empleado en formato JSON
    c.JSON(http.StatusOK, empleado)
}