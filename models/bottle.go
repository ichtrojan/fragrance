package models

import "github.com/jinzhu/gorm"

type Bottle struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(100);unique_index"`
	Slug        string       `gorm:"type:varchar(100);unique_index"`
	Image       string       `gorm:"type:varchar(100)"`
	BottleSizes []BottleSize `gorm:"foreignkey:BottleID"`
}
