package models

import (
	"time"
)

type Comment struct {
	Id        int `gorm:"primaryKey"`
	UserId    int
	User      *User
	PostId    uint
	Post      *Post
	Status    int
	Slug      string `gorm:"type:string;size:255"`
	Title     string `gorm:"type:string;size:200"`
	Content   string
	CreatedAt time.Time
}
