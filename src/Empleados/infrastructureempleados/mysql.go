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
	return empleado, nil
}

func (repo *MySQLRepository) GetAll() ([]entities.Empleado, error) {
	var empleados []entities.Empleado
	rows, err := repo.db.Query("SELECT id, nombre, apellido FROM empleados")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var empleado entities.Empleado
		err := rows.Scan(&empleado.Id, &empleado.Nombre, &empleado.Apellido)
		if err != nil {
			return nil, err
		}
		empleados = append(empleados, empleado)
	}
	return empleados, nil
}


func (repo *MySQLRepository) GetById(id int) (entities.Empleado, error) {
	var empleado entities.Empleado
	err := repo.db.QueryRow("SELECT id, nombre, apellido FROM empleados WHERE id = ?", id).
		Scan(&empleado.Id, &empleado.Nombre, &empleado.Apellido)
	if err != nil {
		return entities.Empleado{}, err
	}
	return empleado, nil
}

func (repo *MySQLRepository) Update(empleado entities.Empleado) (entities.Empleado, error) {
	_, err := repo.db.Exec("UPDATE empleados SET nombre = ?, apellido = ? WHERE id = ?",
		empleado.Nombre, empleado.Apellido, empleado.Id)
	if err != nil {
		return entities.Empleado{}, err
	}
	return empleado, nil
}

func (repo *MySQLRepository) Delete(id int) error {
	_, err := repo.db.Exec("DELETE FROM empleados WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}