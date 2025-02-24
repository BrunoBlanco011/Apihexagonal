package application

import (
	"apiGo/src/Productos/domain"
	"apiGo/src/Productos/domain/entities"
)

type UpdateProductUseCase struct {
	repo domain.ProductRepository
}

func NewUpdateProductUseCase(repo domain.ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{repo: repo}
}

func (uc *UpdateProductUseCase) UpdateProduct(product entities.Product) (entities.Product, error) {
	return uc.repo.Update(product)
}
