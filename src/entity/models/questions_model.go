package models

import "gorm.io/gorm"

type QuestionModel struct {
	gorm.Model
	ID            uint          `gorm:"primaryKey,autoIncrement,not null"`
	Question      string        `gorm:"size:255;not null"`
	Options       []OptionModel `gorm:"foreignKey:QuestionID"`
	CorrectOption uint          `gorm:"not null"`
	Difficulty    string        `gorm:"type:VARCHAR(10);not null;check:difficulty IN ('facil', 'medio', 'dificil')"`
	Points        int           `gorm:"not null"`
}
