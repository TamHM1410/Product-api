package repository

import (
	entity "go-backend-api/internal/product/domain/model/entity"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(product *entity.ProductModel) (*entity.ProductModel, error) {
	toSave := r.db.Create(product)
	if toSave.Error != nil {
		return nil, toSave.Error
	}
	return product, nil
}

func (r *ProductRepo) GetByID(id uint) (*entity.ProductModel, error) {
	var prod entity.ProductModel
	result := r.db.First(&prod, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &prod, nil
}

func (r *ProductRepo) CreateOrUpdate(product *entity.ProductModel) (*entity.ProductModel, error) {

	// Option 1: Sử dụng ID để xác định Create hoặc Update
	if product.ID != 0 {
		// Update existing record
		var existing entity.ProductModel
		result := r.db.First(&existing, product.ID)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return nil, result.Error // ID không tồn tại
			}
			return nil, result.Error
		}

		// Cập nhật tất cả các trường
		updateResult := r.db.Model(&existing).Updates(product)
		if updateResult.Error != nil {
			return nil, updateResult.Error
		}

		// Lấy lại dữ liệu sau khi update
		r.db.First(&existing, product.ID)
		return &existing, nil
	}

	// Create new record
	createResult := r.db.Create(product)

	if createResult.Error != nil {
		return nil, createResult.Error
	}

	return product, nil
}
