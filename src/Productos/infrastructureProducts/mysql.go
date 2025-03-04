package infrastructureproducts

import "database/sql"
import "apiGo/src/Productos/domain/entities"
import "apiGo/src/Productos/domain"

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) domain.ProductRepository {
	return &MySQLRepository{db: db}
}


func (repo *MySQLRepository) Create(producto entities.Product) (entities.Product, error) {
	_, err := repo.db.Exec("INSERT INTO productos (id, nombre, cantidad, precio) VALUES (?, ?,?, ?)",
		producto.Id, producto.Nombre, producto.Cantidad, producto.Precio)
	if err != nil {
		return entities.Product{}, err
	}
	return producto, nil
}

func (repo *MySQLRepository) GetAll() ([]entities.Product, error) {
	rows, err := repo.db.Query("SELECT * FROM productos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []entities.Product{}
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(&product.Id, &product.Nombre, &product.Cantidad, &product.Precio)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *MySQLRepository) GetByID(id int) (*entities.Product, error) {
	var product entities.Product
	err := repo.db.QueryRow("SELECT * FROM productos WHERE id = ?", id).
		Scan(&product.Id, &product.Nombre, &product.Cantidad, &product.Precio)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *MySQLRepository) Update(product entities.Product) (entities.Product, error) {
	_, err := repo.db.Exec("UPDATE productos SET nombre = ?, cantidad = ?, precio = ? WHERE id = ?",
		product.Nombre, product.Cantidad, product.Precio, product.Id)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (repo *MySQLRepository) Delete(id int) error {
	_, err := repo.db.Exec("DELETE FROM productos WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}