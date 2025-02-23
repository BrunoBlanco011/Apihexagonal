package application

import (
	"apiGo/src/Empleados/domain"
	"apiGo/src/Empleados/domain/entities"
)


type ViewAllEmpleadoUseCase struct {
	repo domain.EmpleadoRepository
}

func NewViewAllEmpleadoUseCase(repo domain.EmpleadoRepository) *ViewAllEmpleadoUseCase {
	return &ViewAllEmpleadoUseCase{repo: repo}
}

func (uc *ViewAllEmpleadoUseCase) GetAllEmpleados() ([]entities.Empleado, error) {
	return uc.repo.GetAll()
}
