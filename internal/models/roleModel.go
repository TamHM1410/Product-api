package models
import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;column:name;type:varchar(100);index unique,not null"`
	Description string `gorm:"column:description;type:varchar(100);not null"`
}

func (Role) TableName() string {
	return "role_model"
}