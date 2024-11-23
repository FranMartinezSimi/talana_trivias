package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name  string `gorm:"size:50;not null"`
	Email string `gorm:"size:30;not null;unique"`
}
