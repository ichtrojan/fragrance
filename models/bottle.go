package models

type Bottle struct {
	Name  string `gorm:"type:varchar(100);unique_index"`
	Slug  string `gorm:"type:varchar(100);unique_index"`
	Image string `gorm:"type:varchar(100)"`
}
