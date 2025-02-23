package application

import (
	"apiGo/src/Empleados/domain"
	"apiGo/src/Empleados/domain/entities"
)

type UpdateEmpleadoUseCase struct {
	repo domain.EmpleadoRepository
}

func NewUpdateEmpleadoUseCase(repo domain.EmpleadoRepository) *UpdateEmpleadoUseCase {
	return &UpdateEmpleadoUseCase{repo: repo}
}

func (uc *UpdateEmpleadoUseCase) UpdateEmpleado(empleado entities.Empleado) (entities.Empleado, error) {
	return uc.repo.Update(empleado)
}
