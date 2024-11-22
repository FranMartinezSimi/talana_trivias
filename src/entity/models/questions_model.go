package models

import "gorm.io/gorm"

type QuestionModel struct {
	gorm.Model
	Question string        `gorm:"size:255;not null"`
	Enabled  bool          `gorm:"default:true"`
	TriviaID uint          `gorm:"not null"` // Clave foránea para TriviaModel
	Answer   []AnswerModel `gorm:"foreignKey:QuestionID"`
}
