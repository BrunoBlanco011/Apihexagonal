package main

import (
	"apiGo/src/Empleados/infrastructureempleados"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	EmpleadosDeps, err := infrastructureempleados.NewDependencies()

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

	router.Run(":8080")
}
