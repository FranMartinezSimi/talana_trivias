package models

import (
	"gorm.io/gorm"
	"time"
)

type TriviaQuestionsModel struct {
	gorm.Model
	TriviaID   uint      `gorm:"primaryKey"`
	QuestionID uint      `gorm:"primaryKey"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
