package models

import "gorm.io/gorm"

type AnswerModel struct {
	gorm.Model
	Value      string `gorm:"size:255;not null"`
	IsCorrect  bool   `gorm:"default:false"`
	QuestionID uint   `gorm:"not null"`
	CreatedAt  int64  `gorm:"autoCreateTime"`
}
