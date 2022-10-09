package models

import (
	"time"
)

type User struct {
	ID           int        `gorm:"primarykey;autoIncrement"`
	FirstName    string     `gorm:"type:string;size:100" validate:"required"`
	LastName     string     `gorm:"type:string;size:100" validate:"required"`
	Email        string     `gorm:"type:string;size:100" validate:"required,email"`
	UserName     string     `gorm:"type:string;size:100" validate:"required"`
	PasswordHash string     `gorm:"type:string;size:255" validate:"required"`
	RegisteredAt *time.Time `validate:"required"`
	Posts        []*Post
}
