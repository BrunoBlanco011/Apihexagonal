package controllers

import (
    "net/http"
    "time"

    "apiGo/src/Productos/application"
    "apiGo/src/core"

    "github.com/gin-gonic/gin"
)

type ProductPollingController struct {
    viewAllProductUseCase application.ViewAllProductUseCase
    notifier              *core.UpdateNotifier
}

func NewProductPollingController(useCase application.ViewAllProductUseCase, notifier *core.UpdateNotifier) *ProductPollingController {
    return &ProductPollingController{viewAllProductUseCase: useCase, notifier: notifier}
}

// ShortPolling handles short polling requests
func (controller *ProductPollingController) ShortPolling(c *gin.Context) {
    products, err := controller.viewAllProductUseCase.GetAllProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
    controller.notifier.NotifyUpdate()
}

// LongPolling handles long polling requests
func (controller *ProductPollingController) LongPolling(c *gin.Context) {
    select {
    case <-controller.notifier.Updates:
        println("Se recibió una notificación de cambio")
    case <-time.After(30 * time.Second):
        println("Timeout: No hubo cambios en 30 segundos")
    }

    products, err := controller.viewAllProductUseCase.GetAllProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
}