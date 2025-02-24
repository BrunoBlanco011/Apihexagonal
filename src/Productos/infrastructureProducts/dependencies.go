package infrastructureproducts

import "apiGo/src/Productos/application"
import "apiGo/src/core"

type Dependencies struct {
	AddProductosUseCase *application.AddProductUseCase
	ViewAllProductUseCase *application.ViewAllProductUseCase
	ViewbyIDProductUseCase *application.ViewByIDProductUseCase
	UpdateProductUseCase *application.UpdateProductUseCase
	DeleteProductUseCase *application.DeleteProductUseCase
}

func NewDependencies() (*Dependencies, error) {
	db, err := core.InitDB()
	if err != nil {
		return nil, err
	}

	repo := NewMySQLRepository(db)

	return &Dependencies{
		AddProductosUseCase: application.NewAddProductUseCase(repo),
		ViewAllProductUseCase: application.NewViewAllProductUseCase(repo),
		ViewbyIDProductUseCase: application.NewViewByIDProductUseCase(repo),
		UpdateProductUseCase: application.NewUpdateProductUseCase(repo),
		DeleteProductUseCase: application.NewDeleteProductUseCase(repo),
	}, nil
}
