package product

import "gorm.io/gorm"

type ProductModel struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"uniqueIndex"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2)"`
	Details     string  `gorm:"column:details;type:text"`
	Height      float64 `gorm:"type:decimal(10,2)"`
	Width       float64 `gorm:"type:decimal(10,2)"`
	Length      float64 `gorm:"type:decimal(10,2)"`
	Depth       float64 `gorm:"type:decimal(10,2)"`
}

func (ProductModel) TableName() string {
	return "product_model"
}
