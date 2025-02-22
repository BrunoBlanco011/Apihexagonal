package infrastructureempleados

import (
	"apiGo/src/Empleados/domain"
	"apiGo/src/Empleados/domain/entities"
	"database/sql"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) domain.EmpleadoRepository {
	return &MySQLRepository{db: db}
}

func (repo *MySQLRepository) Create(empleado entities.Empleado) (entities.Empleado, error) {
	_, err := repo.db.Exec("INSERT INTO empleados (id, nombre, apellido) VALUES (?, ?, ?)",
		empleado.Id, empleado.Nombre, empleado.Apellido)
	if err != nil {
		return entities.Empleado{}, err
	}
	return empleado,nil
}