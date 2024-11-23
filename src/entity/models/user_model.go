package models

type UserModel struct {
	ID    uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name  string `gorm:"size:50;not null"`
	Email string `gorm:"size:30;not null;unique"`
}
