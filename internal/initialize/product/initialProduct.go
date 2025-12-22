package initialize

import (
	"go-backend-api/internal/product/application/service"
	product "go-backend-api/internal/product/controller/http"
	"go-backend-api/internal/product/domain/repository"

	"gorm.io/gorm"
)

func InitialProduct(db *gorm.DB) *product.ProductHandler {
	repo := repository.NewProductRepo(db)
	service := service.NewProductService(repo)
	handler := product.NewProductHandler(service)

	return handler

}
