package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	UserID uint    `json:"user_id"`
	User   User    `gorm:"foreignKey:UserID"`
}