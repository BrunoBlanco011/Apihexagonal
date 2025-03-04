package application

import (
    "apiGo/src/Productos/domain"
    "apiGo/src/Productos/domain/entities"
    "apiGo/src/core"
)

type UpdateProductUseCase struct {
    repo     domain.ProductRepository
    notifier *core.UpdateNotifier
}

func NewUpdateProductUseCase(repo domain.ProductRepository, notifier *core.UpdateNotifier) *UpdateProductUseCase {
    return &UpdateProductUseCase{repo: repo, notifier: notifier}
}

func (uc *UpdateProductUseCase) UpdateProduct(product entities.Product) (entities.Product, error) {
    result, err := uc.repo.Update(product)
    if err != nil {
        return entities.Product{}, err
    }
    uc.notifier.NotifyUpdate()
    return result, nil
}