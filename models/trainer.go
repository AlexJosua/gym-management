package models

import "gorm.io/gorm"

type Trainer struct {
	gorm.Model
	Name      string `json:"name" binding:"required"`
	Expertise string `json:"expertise"`
}
