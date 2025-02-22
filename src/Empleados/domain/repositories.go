package domain

import "apiGo/src/Empleados/domain/entities"

type EmpleadoRepository interface {
	//GetAll() ([]entities.Empleado, error)
	//GetById(id int) (entities.Empleado, error)
	Create(empleado entities.Empleado) (entities.Empleado, error)
	//Update(empleado entities.Empleado) (entities.Empleado, error)
	//Delete(id int)error
}