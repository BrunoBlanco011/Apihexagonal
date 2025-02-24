package domain

import "apiGo/src/Productos/domain/entities"

type ProductRepository interface {
	GetAll() ([]entities.Product, error)
	GetByID(id int) (*entities.Product, error)
	Create(product entities.Product) (entities.Product, error)
	Update(product entities.Product) (entities.Product, error)
	Delete(id int) error
}