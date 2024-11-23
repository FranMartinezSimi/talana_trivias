package models

type Ranking struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null"`
	Score  uint `gorm:"not null"`
}
