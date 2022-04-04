package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	UserID      uint   `gorm:"not null" json:"user_id" form:"user_id"`
	ProjectName string `gorm:"not null" json:"project_name" form:"project_name"`
	Task        []Task `gorm:"foreignKey:UserID;references:ID"`
}
