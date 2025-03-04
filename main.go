package main

import (
	"apiGo/src/Empleados/infrastructureempleados"
	"apiGo/src/Productos/infrastructureProducts"
	"apiGo/src/core"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	EmpleadosDeps, err := infrastructureempleados.NewDependencies()

	ProductsDeps, err := infrastructureproducts.NewDependencies()

	if err != nil {
		log.Fatalf("Error initializing dependencies: %v", err)
	}

	notifier := core.NewUpdateNotifier()

	router := gin.Default()


	infrastructureempleados.RegisterEmpleadosRoutes(
		router,
		EmpleadosDeps.AddEmpleadosUseCase,
		EmpleadosDeps.ViewEmpleadoIDUseCase,
		EmpleadosDeps.ViewAllEmpleadoUseCase,
		EmpleadosDeps.UpdateEmpleadoUseCase,
		EmpleadosDeps.DeleteEmpleadoUseCase,
		notifier,
	)

	infrastructureproducts.RegisterProductosRoutes(
		router,
		ProductsDeps.AddProductosUseCase,
		ProductsDeps.ViewAllProductUseCase,
		ProductsDeps.ViewbyIDProductUseCase,
		ProductsDeps.UpdateProductUseCase,
		ProductsDeps.DeleteProductUseCase,
		notifier,
	)
	router.Run(":8080")
}
