package entities

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID     uint   `gorm:"not null" json:"user_id" form:"user_id"`
	ProjectID  *uint  `json:"project_id" form:"project_id"`
	TaskName   string `gorm:"not null" json:"task_name" form:"task_name"`
	TaskStatus string `gorm:"not null;default:not completed" json:"task_status" form:"task_status"`
}
