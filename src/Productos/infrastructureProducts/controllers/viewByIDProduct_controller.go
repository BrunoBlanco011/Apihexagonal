package controllers

import (
	"apiGo/src/Productos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type viewByIDProductController struct {
	useCase application.ViewByIDProductUseCase
}

func NewViewByIDProductController(useCase application.ViewByIDProductUseCase) *viewByIDProductController {
	return &viewByIDProductController{useCase: useCase}
}

func (controller *viewByIDProductController) ViewbyIDProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	product, err := controller.useCase.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	c.JSON(http.StatusOK, product)
}
