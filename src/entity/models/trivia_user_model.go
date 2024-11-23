package models

type TriviaUser struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint `gorm:"not null"`
	TriviaID uint `gorm:"not null"`
}
