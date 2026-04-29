package models

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	Title   string    `json:"title"`
	Price   float64   `json:"price"`
	DueDate time.Time `json:"due_date"`
	UserID  uint      `json:"user_id"`
	User    User      `gorm:"foreignKey:UserID"`
}