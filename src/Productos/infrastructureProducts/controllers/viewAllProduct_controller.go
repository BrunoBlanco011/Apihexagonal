package controllers

import (
	"apiGo/src/Productos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type viewAllProductController struct {
	useCase application.ViewAllProductUseCase
}

func NewViewAllProductController(useCase application.ViewAllProductUseCase) *viewAllProductController {
	return &viewAllProductController{useCase: useCase}
}

func (controller *viewAllProductController) ViewAllProduct(c *gin.Context) {
	result, err := controller.useCase.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
