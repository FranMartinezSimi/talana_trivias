package models

import "gorm.io/gorm"

type ScoreModel struct {
	gorm.Model
	UserID      uint `gorm:"not null"`
	TriviaID    uint `gorm:"not null"`
	Score       uint `gorm:"default:0"`
	IsCompleted bool `gorm:"default:false"`
}
