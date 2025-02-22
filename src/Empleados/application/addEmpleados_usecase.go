package application

import (
	"apiGo/src/Empleados/domain"
	"apiGo/src/Empleados/domain/entities"
)

type AddEmpleadosUseCase struct {
	repo domain.EmpleadoRepository
}

func NewAddEmpleadosUseCase(repo domain.EmpleadoRepository) *AddEmpleadosUseCase {
	return &AddEmpleadosUseCase{repo: repo}
}

func (uc *AddEmpleadosUseCase) AddEmpleado(empleado entities.Empleado) (entities.Empleado, error) {
	return uc.repo.Create(empleado)
}
