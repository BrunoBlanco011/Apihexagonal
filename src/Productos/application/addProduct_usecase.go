package application

import (
	"apiGo/src/Productos/domain"
	"apiGo/src/Productos/domain/entities"
)

type AddProductUseCase struct {
	repo domain.ProductRepository

}

func NewAddProductUseCase(repo domain.ProductRepository) *AddProductUseCase {
	return &AddProductUseCase{repo: repo}
}

func (uc *AddProductUseCase) AddProduct(product entities.Product) (entities.Product, error) {
	return uc.repo.Create(product)
}