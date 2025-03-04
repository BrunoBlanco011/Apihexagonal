package application

import (
    "apiGo/src/Empleados/domain"
    "apiGo/src/Empleados/domain/entities"
    "apiGo/src/core"
)

type UpdateEmpleadoUseCase struct {
    repo     domain.EmpleadoRepository
    notifier *core.UpdateNotifier
}

func NewUpdateEmpleadoUseCase(repo domain.EmpleadoRepository, notifier *core.UpdateNotifier) *UpdateEmpleadoUseCase {
    return &UpdateEmpleadoUseCase{repo: repo, notifier: notifier}
}

func (uc *UpdateEmpleadoUseCase) UpdateEmpleado(empleado entities.Empleado) (entities.Empleado, error) {
    result, err := uc.repo.Update(empleado)
    if err != nil {
        return entities.Empleado{}, err
    }
    uc.notifier.NotifyUpdate()
    return result, nil
}