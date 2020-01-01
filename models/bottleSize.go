package models

import (
	"github.com/jinzhu/gorm"
)

type BottleSize struct {
	gorm.Model
	BottleID uint
	Size     string  `gorm:"type:varchar(100)"`
	Price    float64 `gorm:"type:decimal(10,2);not null"`
}
