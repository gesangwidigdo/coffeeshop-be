package models

import "gorm.io/gorm"

type Outlet struct {
	gorm.Model
	Outlet_name string     `gorm:"not null; type:varchar(50)"`
	Address     string     `gorm:"not null; type:text"`
	City        string     `gorm:"not null; type:varchar(150)"`
	Employees   []Employee `gorm:"foreignKey:Outlet_id" json:"-"`
}
