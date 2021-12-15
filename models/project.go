package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255);not null" json:"name" form:"name"`
	Tasks  []Task
	UserID uint
}
