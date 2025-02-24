package controllers

import (
	"apiGo/src/Productos/application"
	"apiGo/src/Productos/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	useCase application.UpdateProductUseCase
}

func NewUpdateProductController(useCase application.UpdateProductUseCase) *UpdateProductController {
	return &UpdateProductController{useCase: useCase}
}

func (controller *UpdateProductController) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var product entities.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product.Id = id
	result, err := controller.useCase.UpdateProduct(product)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}
