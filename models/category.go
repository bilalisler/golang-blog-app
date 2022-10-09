package models

type Category struct {
	Id     int    `gorm:"primaryKey"`
	Name   string `gorm:"type:string;size:200"`
	Parent int
	Slug   string `gorm:"type:string;size:255"`
}
