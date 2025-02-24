package infrastructureproducts

import "github.com/gin-gonic/gin"
import "apiGo/src/Productos/application"
import "apiGo/src/Productos/infrastructureproducts/controllers"

func RegisterProductosRoutes(
	router *gin.Engine,
	addProductosUseCase *application.AddProductUseCase,
	viewAllProductosUseCase *application.ViewAllProductUseCase,
	viewbyIDProductosUseCase *application.ViewByIDProductUseCase,
	updateProductosUseCase *application.UpdateProductUseCase,
	deleteProductosUseCase *application.DeleteProductUseCase,
) {
	addProductController := controllers.NewAddProductController(*addProductosUseCase)
	viewAllProductController := controllers.NewViewAllProductController(*viewAllProductosUseCase)
	viewByIDProductController := controllers.NewViewByIDProductController(*viewbyIDProductosUseCase)
	updateProductController := controllers.NewUpdateProductController(*updateProductosUseCase)
	deleteProductController := controllers.NewDeleteProductController(*deleteProductosUseCase)

	router.POST("/productos", addProductController.AddProduct)
	router.GET("/productos", viewAllProductController.ViewAllProduct)
	router.GET("/productos/:id", viewByIDProductController.ViewbyIDProduct)
	router.PUT("/productos/:id", updateProductController.UpdateProduct)
	router.DELETE("/productos/:id", deleteProductController.DeleteProduct)
	
}