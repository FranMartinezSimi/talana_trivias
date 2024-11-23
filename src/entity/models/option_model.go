package models

type Option struct {
	ID         uint   `gorm:"primaryKey"`
	Text       string `gorm:"not null"`
	QuestionID uint   `gorm:"not null"`
}
