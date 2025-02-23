package application

import "apiGo/src/Empleados/domain"

type DeleteEmpleadoUseCase struct {
	repo domain.EmpleadoRepository
}

func NewDeleteEmpleadoUseCase(repo domain.EmpleadoRepository) *DeleteEmpleadoUseCase {
	return &DeleteEmpleadoUseCase{repo: repo}
}

func (uc *DeleteEmpleadoUseCase) DeleteEmpleado(id int) error {
	return uc.repo.Delete(id)
}