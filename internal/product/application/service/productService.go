package service

import (
	"go-backend-api/internal/product/application/service/dto"
	product "go-backend-api/internal/product/domain/model/entity"
	"go-backend-api/pb"
)

type ProductService interface {
	CreateProduct(payload *dto.ProductDTO) (*product.ProductModel, error)

	GetProductByID(id uint) (*product.ProductModel, error)

	CreateOrUpdate(payload *pb.ProductSimpleDTO) (*product.ProductModel, error)

	// createOrUpdate(payload *dto.ProductDTO) (*product.ProductModel, error)
	// UpdateProduct(id uint, payload *dto.ProductDTO) (*dto.ProductDTO, error)
	// DeleteProduct(id uint) error
	// ListProducts() ([]*dto.ProductDTO, error)
}
