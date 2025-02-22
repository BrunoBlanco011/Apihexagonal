package domain

import (
	"api-hexagonal/src/Empleados/domain"
)


type EmpleadoRepository interface {
	GetAll() ([]Empleado, error)
	GetById(id int) (Empleado, error)
	Create(empleado Empleado) (Empleado, error)
	Update(empleado Empleado) (Empleado, error)
	Delete(id int) error
}
