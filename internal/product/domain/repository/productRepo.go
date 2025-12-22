package repository

import (
	product "go-backend-api/internal/product/domain/model/entity"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(product *product.ProductModel) (*product.ProductModel, error) {
	toSave := r.db.Create(product)
	if toSave.Error != nil {
		return nil, toSave.Error
	}
	return product, nil
}

func (r *ProductRepo) GetByID(id uint) (*product.ProductModel, error) {
	var prod product.ProductModel
	result := r.db.First(&prod, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &prod, nil
}
