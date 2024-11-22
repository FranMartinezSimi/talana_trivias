package models

type Participation struct {
	ID       uint     `gorm:"primaryKey"`
	UserID   uint     `gorm:"not null"`
	TriviaID uint     `gorm:"not null"`
	Score    int      `gorm:"default:0"`
	Answers  []Answer `gorm:"foreignKey:ParticipationID"`
}
