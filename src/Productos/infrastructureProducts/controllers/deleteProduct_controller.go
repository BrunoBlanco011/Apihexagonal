package controllers

import (
	"apiGo/src/Productos/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCase application.DeleteProductUseCase
}

func NewDeleteProductController(useCase application.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{useCase: useCase}
}

func (controller *DeleteProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = controller.useCase.DeleteProduct(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Producto eliminado"})
}
