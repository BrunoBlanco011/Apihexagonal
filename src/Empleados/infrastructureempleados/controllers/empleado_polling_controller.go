package controllers

import (
    "net/http"
    "time"

    "apiGo/src/Empleados/application"
    "apiGo/src/core"

    "github.com/gin-gonic/gin"
)

type EmpleadoPollingController struct {
    viewAllEmpleadoUseCase application.ViewAllEmpleadoUseCase
    notifier               *core.UpdateNotifier
}

func NewEmpleadoPollingController(useCase application.ViewAllEmpleadoUseCase, notifier *core.UpdateNotifier) *EmpleadoPollingController {
    return &EmpleadoPollingController{viewAllEmpleadoUseCase: useCase, notifier: notifier}
}

// ShortPolling handles short polling requests
func (controller *EmpleadoPollingController) ShortPolling(c *gin.Context) {
    empleados, err := controller.viewAllEmpleadoUseCase.GetAllEmpleados()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, empleados)
    controller.notifier.NotifyUpdate()
}

// LongPolling handles long polling requests
func (controller *EmpleadoPollingController) LongPolling(c *gin.Context) {
    select {
    case <-controller.notifier.Updates:
        println("Se recibió una notificación de cambio")
    case <-time.After(30 * time.Second):
        println("Timeout: No hubo cambios en 30 segundos")
    }

    empleados, err := controller.viewAllEmpleadoUseCase.GetAllEmpleados()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, empleados)
}