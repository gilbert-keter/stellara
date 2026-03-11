package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	first_name string `gorm:"type:varchar(255)"`
	last_name  string `gorm:"type:varchar(255)"`
	age        string `gorm:"type:int"`
}
