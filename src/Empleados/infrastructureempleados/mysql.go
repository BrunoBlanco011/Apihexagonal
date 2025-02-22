package infrastructureempleados

import (
	"database/sql"
	"api-hexagonal/src/Empleados/domain/entities"
	"api-hexagonal/src/Empleados/domain"

)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) domain.EmpleadoRepository {
	return &MySQLRepository{db: db}
}

func (repo *MySQLRepository) Create(empleado entities.Empleado) (entities.Empleado, error) {
	_, err := repo.db.Exec("INSERT INTO empleados (id, nombre, apellido,) VALUES (?, ?, ?)", empleado.id, empleado.Nombre, empleado.Apellido,)
	if err != nil {	
		return entities.Empleado{}, err
	}
	return empleado, nil
}