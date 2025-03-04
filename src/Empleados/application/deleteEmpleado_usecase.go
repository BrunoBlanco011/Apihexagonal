package application

import (
    "apiGo/src/Empleados/domain"
    "apiGo/src/core"
)

type DeleteEmpleadoUseCase struct {
    repo     domain.EmpleadoRepository
    notifier *core.UpdateNotifier
}

func NewDeleteEmpleadoUseCase(repo domain.EmpleadoRepository, notifier *core.UpdateNotifier) *DeleteEmpleadoUseCase {
    return &DeleteEmpleadoUseCase{repo: repo, notifier: notifier}
}

func (uc *DeleteEmpleadoUseCase) DeleteEmpleado(id int) error {
    err := uc.repo.Delete(id)
    if err != nil {
        return err
    }
    uc.notifier.NotifyUpdate()
    return nil
}