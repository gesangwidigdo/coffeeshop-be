package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Ingredient_name string `gorm:"not null; type:varchar(100)"`
	Stock           int    `gorm:"not null; type:integer; default:0"`
}
