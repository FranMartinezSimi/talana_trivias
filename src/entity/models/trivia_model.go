package models

type Trivia struct {
	ID          uint        `gorm:"primaryKey"`
	Name        string      `gorm:"not null"`
	Description string      `gorm:"not null"`
	Questions   []Question  `gorm:"many2many:trivia_questions;"`
	Users       []UserModel `gorm:"many2many:trivia_users;"`
}
