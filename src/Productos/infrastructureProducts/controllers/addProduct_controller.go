package controllers

import (
	"apiGo/src/Productos/application"
	"apiGo/src/Productos/domain/entities"

	"github.com/gin-gonic/gin"
)

type AddProductController struct {
	useCase application.AddProductUseCase
}

func NewAddProductController(useCase application.AddProductUseCase) *AddProductController {
	return &AddProductController{useCase: useCase}
}

func (controller *AddProductController) AddProduct(c *gin.Context) {
	var product entities.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.useCase.AddProduct(product)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}
