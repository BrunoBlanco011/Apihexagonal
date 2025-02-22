package infrastructureempleados

import ( 
"api-hexagonal/src/Empleados/application"
"api-hexagonal/src/Empleados/infrastructureempleados/controllers"

"github.com/gin-gonic/gin"
)

func RegisterEmpleadosRoutes(
	router *gin.Engine,
	addEmpleadosUseCase *application.AddEmpleadosUseCase,
) {
	addEmpleadosController := controllers.NewAddEmpleadosController(*addEmpleadosUseCase)
	router.POST("/empleados", addEmpleadosController.AddEmpleado)
}