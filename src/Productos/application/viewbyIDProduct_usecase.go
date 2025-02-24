package application

import (
	"apiGo/src/Productos/domain"
	"apiGo/src/Productos/domain/entities"
)

type ViewByIDProductUseCase struct {
	repo domain.ProductRepository
}

func NewViewByIDProductUseCase(repo domain.ProductRepository) *ViewByIDProductUseCase {
	return &ViewByIDProductUseCase{repo: repo}
}

func (uc *ViewByIDProductUseCase) GetProduct(id int) (*entities.Product, error) {
	return uc.repo.GetByID(id)

}
