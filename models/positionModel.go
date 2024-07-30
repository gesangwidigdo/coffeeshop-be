package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Position_name string     `gorm:"not null; type:varchar(20)"`
	Employees     []Employee `gorm:"foreignKey:Position_id" json:"-"`
}
