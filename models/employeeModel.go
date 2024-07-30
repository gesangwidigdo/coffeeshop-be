package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Employee_name    string   `gorm:"not null; type:varchar(150)"`
	Telephone_number string   `gorm:"not null; type:varchar(15)"`
	Username         string   `gorm:"not null; type:varchar(50)"`
	Password         string   `gorm:"not null; type: varchar(50)"`
	Address          string   `gorm:"not null; type:varchar(255)"`
	Gender           string   `gorm:"not null; type:varchar(10)"`
	Position_id      uint     `gorm:"index"`
	Outlet_id        uint     `gorm:"index"`
	Position         Position `gorm:"foreignKey:Position_id;references:ID"`
	Outlet           Outlet   `gorm:"foreignKey:Outlet_id;references:ID"`
}
