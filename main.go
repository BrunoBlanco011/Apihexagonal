package main

import (
	"apiGo/src/Empleados/infrastructureempleados"
	"apiGo/src/Productos/infrastructureProducts"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	EmpleadosDeps, err := infrastructureempleados.NewDependencies()

	ProductsDeps, err := infrastructureproducts.NewDependencies()

	if err != nil {
		log.Fatalf("Error initializing dependencies: %v", err)
	}

	router := gin.Default()

	infrastructureempleados.RegisterEmpleadosRoutes(
		router,
		EmpleadosDeps.AddEmpleadosUseCase,
		EmpleadosDeps.ViewEmpleadoIDUseCase,
		EmpleadosDeps.ViewAllEmpleadoUseCase,
		EmpleadosDeps.UpdateEmpleadoUseCase,
		EmpleadosDeps.DeleteEmpleadoUseCase,
	)

	infrastructureproducts.RegisterProductosRoutes(
		router,
		ProductsDeps.AddProductosUseCase,
		ProductsDeps.ViewAllProductUseCase,
		ProductsDeps.ViewbyIDProductUseCase,
		ProductsDeps.UpdateProductUseCase,
		ProductsDeps.DeleteProductUseCase,
	)
	router.Run(":8080")
}
