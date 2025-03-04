package infrastructureproducts

import (
	"apiGo/src/Productos/application"
	"apiGo/src/Productos/infrastructureProducts/controllers"
	"apiGo/src/core"

	"github.com/gin-gonic/gin"
)

func RegisterProductosRoutes(
	router *gin.Engine,
	addProductosUseCase *application.AddProductUseCase,
	viewAllProductosUseCase *application.ViewAllProductUseCase,
	viewbyIDProductosUseCase *application.ViewByIDProductUseCase,
	updateProductosUseCase *application.UpdateProductUseCase,
	deleteProductosUseCase *application.DeleteProductUseCase,
	notifier *core.UpdateNotifier,
) {
	addProductController := controllers.NewAddProductController(*addProductosUseCase)
	viewAllProductController := controllers.NewViewAllProductController(*viewAllProductosUseCase)
	viewByIDProductController := controllers.NewViewByIDProductController(*viewbyIDProductosUseCase)
	updateProductController := controllers.NewUpdateProductController(*updateProductosUseCase)
	deleteProductController := controllers.NewDeleteProductController(*deleteProductosUseCase)
	productPollingController := controllers.NewProductPollingController(*viewAllProductosUseCase, notifier)

	router.POST("/productos", addProductController.AddProduct)
	router.GET("/productos", viewAllProductController.ViewAllProduct)
	router.GET("/productos/:id", viewByIDProductController.ViewbyIDProduct)
	router.PUT("/productos/:id", updateProductController.UpdateProduct)
	router.DELETE("/productos/:id", deleteProductController.DeleteProduct)
	router.GET("/productos/shortpolling", productPollingController.ShortPolling)
	router.GET("/productos/longpolling", productPollingController.LongPolling)

}
