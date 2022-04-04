package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string    `gorm:"not null" json:"name" form:"name"`
	Email      string    `gorm:"unique;not null" json:"email" form:"email"`
	Password   string    `gorm:"not null" json:"password" form:"password"`
	Gender     string    `gorm:"not null" json:"gender" form:"gender"`
	Age        uint      `gorm:"not null" json:"age" form:"age"`
	Profession string    `gorm:"not null" json:"profession" form:"profession"`
	Task       []Task    `gorm:"foreignKey:UserID;references:ID"`
	Project    []Project `gorm:"foreignKey:UserID;references:ID"`
}
