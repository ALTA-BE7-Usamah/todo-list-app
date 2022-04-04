package entities

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID     string    `gorm:"not null" json:"user_id" form:"user_id"`
	ProjectID  string    `json:"project_id" form:"project_id"`
	TaskName   string    `gorm:"not null" json:"task_name" form:"task_name"`
	DueDate    time.Time `gorm:"not null" json:"due_date" form:"due_date"`
	TaskStatus string    `gorm:"not null;default:not completed" json:"task_status" form:"task_status"`
	Project    Project   `gorm:"foreignKey:ProjectID;references:ID"`
}
