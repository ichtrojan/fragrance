package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);unique_index"`
	Slug  string `gorm:"type:varchar(100);not null;unique_index"`
	Image string `gorm:"type:varchar(100);not null;unique_index"`
}
