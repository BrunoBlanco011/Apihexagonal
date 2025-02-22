package infrastructureempleados

import(
	"api-hexagonal/src/core"
	"api-hexagonal/src/Empleados/application"
)

type Dependencies struct {
		AddEmpleadosUseCase *application.AddEmpleadosUseCase
	}

func NewDependencies() (*Dependencies, error) {
	db, err := core.InitDB()
	if err != nil {
		return nil, err
	}
	
	repo := NewMySQLRepository(db)

	return &Dependencies{
		AddEmpleadosUseCase: application.NewAddEmpleadosUseCase(repo),
	}, nil
}