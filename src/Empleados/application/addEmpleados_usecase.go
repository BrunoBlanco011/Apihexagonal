package application

import (
    "apiGo/src/Empleados/domain"
    "apiGo/src/Empleados/domain/entities"
    "apiGo/src/core"
)

type AddEmpleadosUseCase struct {
    repo     domain.EmpleadoRepository
    notifier *core.UpdateNotifier
}

func NewAddEmpleadosUseCase(repo domain.EmpleadoRepository, notifier *core.UpdateNotifier) *AddEmpleadosUseCase {
    return &AddEmpleadosUseCase{repo: repo, notifier: notifier}
}

func (uc *AddEmpleadosUseCase) AddEmpleado(empleado entities.Empleado) (entities.Empleado, error) {
    result, err := uc.repo.Create(empleado)
    if err != nil {
        return entities.Empleado{}, err
    }
    uc.notifier.NotifyUpdate()
    return result, nil
}