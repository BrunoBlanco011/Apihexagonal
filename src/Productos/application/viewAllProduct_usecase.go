package application

import (
	"apiGo/src/Productos/domain"
	"apiGo/src/Productos/domain/entities"
)

type ViewAllProductUseCase struct {
	repo domain.ProductRepository
}

func NewViewAllProductUseCase(repo domain.ProductRepository) *ViewAllProductUseCase {
	return &ViewAllProductUseCase{repo: repo}
}

func (uc *ViewAllProductUseCase) GetAllProducts() ([]entities.Product, error) {
	return uc.repo.GetAll()
}
