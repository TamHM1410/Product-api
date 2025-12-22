package product

import "gorm.io/gorm"

type ProductModel struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"uniqueIndex"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2)"`
	Stock       int     `gorm:"type:int"`
	Details     string  `gorm:"column:details;type:text"`
}

func (ProductModel) TableName() string {
	return "product_model"
}
