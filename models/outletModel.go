package models

import "gorm.io/gorm"

type Outlet struct {
	gorm.Model
	Outlet_name string
	Address string
	City string
}