package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`
	Age      int
	UUID     string `gorm:"uniqueIndex;column:uuid;type:char(255);index unique,not null"`
}

func (User) TableName() string {
	return "user_model"
}
