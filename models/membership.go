package models

import "gorm.io/gorm"

type Membership struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	UserID      uint    `json:"user_id"`
	User        User    `gorm:"foreignKey:UserID"`
}
