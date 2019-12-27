package models

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"not null"`
}
