package model

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	Name string  `gorm:"type:VARCHAR(100); NOT NULL;" json:"department_name"`
}
