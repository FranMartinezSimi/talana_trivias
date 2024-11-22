package models

import (
	"gorm.io/gorm"
	"time"
)

type TriviaModel struct {
	gorm.Model
	Title     string          `gorm:"size:50;not null"`
	Enabled   bool            `gorm:"default:true"`
	UserID    uint            `gorm:"not null"`
	Questions []QuestionModel `gorm:"foreignKey:TriviaID"` // Uno a muchos
	CreatedAt time.Time       `gorm:"autoCreateTime"`
}
