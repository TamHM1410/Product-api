package service

import (
	"go-backend-api/internal/product/application/service/dto"
	product "go-backend-api/internal/product/domain/model/entity"
	"go-backend-api/internal/product/domain/repository"
)

type productService struct {
	productRepo *repository.ProductRepo
}

func (ps *productService) CreateProduct(payload *dto.ProductDTO) (*product.ProductModel, error) {
	var toSave product.ProductModel

	toSave.Name = payload.Name
	toSave.Description = payload.Description
	toSave.Price = payload.Price

	createdProduct, err := ps.productRepo.Create(&toSave)

	return createdProduct, err
}
func (ps *productService) GetProductByID(id uint) (*product.ProductModel, error) {
	// Implementation for getting a product by ID
	product, err := ps.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func NewProductService(repo *repository.ProductRepo) ProductService {
	return &productService{
		productRepo: repo,
	}
}
