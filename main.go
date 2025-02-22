
package main
import (
	"log"
	"APIHEXAGONAL/api-hexagonal/src/Empleados/infrastructureempleados"
	
	

	"github.com/gin-gonic/gin"
)

func main() {
	
	EmpleadosDeps, err := infrastructureempleados.NewDependencies()

	if err != nil {
		log.Fatalf("Error initializing dependencies: %v", err)
	}

	router := gin.Default()

	infrastructureempleados.RegisterEmpleadosRoutes(router, EmpleadosDeps.AddEmpleadosUseCase)

	router.Run(":8080")
}