package application

import (
    "apiGo/src/Productos/domain"
    "apiGo/src/core"
)

type DeleteProductUseCase struct {
    repo     domain.ProductRepository
    notifier *core.UpdateNotifier
}

func NewDeleteProductUseCase(repo domain.ProductRepository, notifier *core.UpdateNotifier) *DeleteProductUseCase {
    return &DeleteProductUseCase{repo: repo, notifier: notifier}
}

func (uc *DeleteProductUseCase) DeleteProduct(id int) error {
    err := uc.repo.Delete(id)
    if err != nil {
        return err
    }
    uc.notifier.NotifyUpdate()
    return nil
}