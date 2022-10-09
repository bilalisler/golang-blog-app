package models

import (
	"time"
)

type Post struct {
	ID         int `gorm:"primarykey;autoIncrement"`
	UserID     int `validate:"required" json:"user_id"`
	User       *User
	Status     int8   `validate:"required" json:"status"`
	Slug       string `gorm:"type:string;size:255" validate:"required"`
	Title      string `gorm:"type:string;size:200" validate:"required" json:"title"`
	Content    string `validate:"required" json:"content"`
	CategoryId int    `validate:"required" json:"category_id"`
	Category   *Category
	Comments   []*Comment
	CreatedAt  time.Time `validate:"required"`
}
