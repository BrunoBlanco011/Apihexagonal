package infrastructureempleados

import (
	"apiGo/src/Empleados/application"
	"apiGo/src/Empleados/infrastructureempleados/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterEmpleadosRoutes(
	router *gin.Engine,
	addEmpleadosUseCase *application.AddEmpleadosUseCase,
) {
	addEmpleadosController := controllers.NewAddEmpleadosController(*addEmpleadosUseCase)
	router.POST("/empleados", addEmpleadosController.AddEmpleado)
}
