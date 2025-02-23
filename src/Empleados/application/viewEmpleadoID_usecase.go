package application

import (
	"apiGo/src/Empleados/domain"
	"apiGo/src/Empleados/domain/entities"
)

type ViewEmpleadoIDUseCase struct {
	repo domain.EmpleadoRepository // Se define el repositorio
}

// NewViewEmpleadoIDUseCase crea una nueva instancia de ViewEmpleadoIDUseCase
func NewViewEmpleadoIDUseCase(repo domain.EmpleadoRepository) *ViewEmpleadoIDUseCase {
	return &ViewEmpleadoIDUseCase{repo: repo}
}

// GetEmpleado obtiene un empleado por su id
func (uc *ViewEmpleadoIDUseCase) GetEmpleado(id int) (entities.Empleado, error) {
	return uc.repo.GetById(id)
}