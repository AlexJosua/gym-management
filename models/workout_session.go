package models

import "gorm.io/gorm"

type WorkoutSession struct {
	gorm.Model
	UserID    uint    `json:"user_id"`    // peserta
	TrainerID uint    `json:"trainer_id"` // pembimbing
	Date      string  `json:"date"`
	Duration  int     `json:"duration"`
	User      User    `gorm:"foreignKey:UserID"`
	Trainer   Trainer `gorm:"foreignKey:TrainerID"`
}
