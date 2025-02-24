package application

import "apiGo/src/Productos/domain"

type DeleteProductUseCase struct {
	repo domain.ProductRepository
}

func NewDeleteProductUseCase(repo domain.ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{repo: repo}
}

func (uc *DeleteProductUseCase) DeleteProduct(id int) error {
	return uc.repo.Delete(id)
}
