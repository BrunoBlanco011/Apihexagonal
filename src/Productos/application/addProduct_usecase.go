package application

import (
    "apiGo/src/Productos/domain"
    "apiGo/src/Productos/domain/entities"
    "apiGo/src/core"
)

type AddProductUseCase struct {
    repo     domain.ProductRepository
    notifier *core.UpdateNotifier
}

func NewAddProductUseCase(repo domain.ProductRepository, notifier *core.UpdateNotifier) *AddProductUseCase {
    return &AddProductUseCase{repo: repo, notifier: notifier}
}

func (uc *AddProductUseCase) AddProduct(product entities.Product) (entities.Product, error) {
    result, err := uc.repo.Create(product)
    if err != nil {
        return entities.Product{}, err
    }
    uc.notifier.NotifyUpdate()
    return result, nil
}