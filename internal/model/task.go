package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	Completed   bool   `gorm:"default:false"`
}
