package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	UserID        string `gorm:"not null" json:"user_id" form:"user_id"`
	ProjectName   string `gorm:"not null" json:"project_name" form:"project_name"`
	ProjectStatus string `gorm:"not null;default:not completed" json:"project_status" form:"project_status"`
	Task          []Task `gorm:"foreignKey:UserID;references:ID"`
}
