package infrastructureempleados

import (
	"apiGo/src/Empleados/application"
	"apiGo/src/Empleados/infrastructureempleados/controllers"
	"apiGo/src/core"

	"github.com/gin-gonic/gin"
)

func RegisterEmpleadosRoutes(
	router *gin.Engine,
	addEmpleadosUseCase *application.AddEmpleadosUseCase,
	viewEmpleadoIDUseCase *application.ViewEmpleadoIDUseCase,
	viewAllEmpleadoUseCase *application.ViewAllEmpleadoUseCase,
	updateEmpleadosUseCase *application.UpdateEmpleadoUseCase,
	deleteEmpleadoUseCase *application.DeleteEmpleadoUseCase,
	notifier *core.UpdateNotifier,
) {
	addEmpleadosController := controllers.NewAddEmpleadosController(*addEmpleadosUseCase)
	viewEmpleadoIDController := controllers.NewViewEmpleadoIDController(*viewEmpleadoIDUseCase)
	viewAllEmpleadoController := controllers.NewViewAllEmpleadoController(*viewAllEmpleadoUseCase)
	updateEmpleadoController := controllers.NewUpdateEmpleadoController(*updateEmpleadosUseCase)
	deleteEmpleadoController := controllers.NewDeleteEmpleadoController(*deleteEmpleadoUseCase)
	empleadoPollingController := controllers.NewEmpleadoPollingController(*viewAllEmpleadoUseCase, notifier)
	

	router.GET("/empleados/:id", viewEmpleadoIDController.ViewEmpleadoID)
	router.POST("/empleados", addEmpleadosController.AddEmpleado)
	router.GET("/empleados", viewAllEmpleadoController.ViewAllEmpleado)
	router.PUT("/empleados/:id", updateEmpleadoController.UpdateEmpleado)
	router.DELETE("/empleados/:id", deleteEmpleadoController.DeleteEmpleado)
	router.GET("/empleados/longpolling", empleadoPollingController.LongPolling)
    router.GET("/empleados/shortpolling", empleadoPollingController.ShortPolling)
}
