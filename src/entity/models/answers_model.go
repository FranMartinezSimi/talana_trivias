package models

type Answer struct {
	ID              uint `gorm:"primaryKey"`
	ParticipationID uint `gorm:"not null"`
	QuestionID      uint `gorm:"not null"`
	SelectedOption  uint `gorm:"not null"`
	IsCorrect       bool `gorm:"not null"`
}
