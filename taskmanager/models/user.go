package models

type User struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Role string `gorm:"not null" json:"role"`
}