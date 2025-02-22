package infrastructureempleados

import (
	"apiGo/src/Empleados/application"
	"apiGo/src/core"
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
