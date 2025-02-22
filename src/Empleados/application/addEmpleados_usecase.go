package application

import (
	"api-hexagonal/src/Empleados/domain/entities"
	"api-hexagonal/src/Empleados/domain"
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