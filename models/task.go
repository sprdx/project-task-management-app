package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string `gorm:"type:mediumtext;not null" json:"description" form:"description"`
	Label       string `gorm:"type:char(50)" json:"label" form:"label"`
	Priority    string `gorm:"type:char(50)" json:"priority" form:"priority"`
	Schedule    time.Time
	Status      string `gorm:"type:enum('not completed', 'completed');default:not completed" json:"status" form:"status"`
	UserID      uint   `json:"user_id" form:"user_id"`
	ProjectID   *uint  `json:"project_id" form:"project_id"`
}
