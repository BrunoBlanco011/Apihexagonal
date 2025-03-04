package infrastructureempleados

import (
	"apiGo/src/Empleados/application"
	"apiGo/src/core"
)

type Dependencies struct {
	AddEmpleadosUseCase *application.AddEmpleadosUseCase
	ViewEmpleadoIDUseCase *application.ViewEmpleadoIDUseCase
	ViewAllEmpleadoUseCase *application.ViewAllEmpleadoUseCase
	UpdateEmpleadoUseCase  *application.UpdateEmpleadoUseCase
	DeleteEmpleadoUseCase *application.DeleteEmpleadoUseCase
}

func NewDependencies() (*Dependencies, error) {
	db, err := core.InitDB()
	if err != nil {
		return nil, err
	}

	repo := NewMySQLRepository(db)
	notifier := core.NewUpdateNotifier()

	return &Dependencies{
		AddEmpleadosUseCase: application.NewAddEmpleadosUseCase(repo, notifier),
		ViewEmpleadoIDUseCase: application.NewViewEmpleadoIDUseCase(repo),
		ViewAllEmpleadoUseCase: application.NewViewAllEmpleadoUseCase(repo),
		UpdateEmpleadoUseCase: application.NewUpdateEmpleadoUseCase(repo, notifier),
		DeleteEmpleadoUseCase: application.NewDeleteEmpleadoUseCase(repo, notifier),
		
	}, nil
}
