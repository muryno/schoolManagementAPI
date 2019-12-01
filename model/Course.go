package model

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Name string `gorm:"type:VARCHAR(100); NOT NULL;" json:"department_name"`
	Code string `gorm:"not null" json:"code"`
}
