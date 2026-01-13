package product

import "gorm.io/gorm"

type ProductVariantModel struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"type:varchar(100)"`
	HexName string `gorm:"type:varchar(100)"`
}

func (ProductVariantModel) TableName() string {
	return "product_variant_model"
}
