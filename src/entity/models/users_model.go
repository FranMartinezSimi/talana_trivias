package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name     string        `gorm:"size:50;not null"`
	Email    string        `gorm:"size:30;not null;unique"`
	Password string        `gorm:"size:50;not null"`
	Trivia   []TriviaModel `gorm:"foreignKey:UserID"`
}
