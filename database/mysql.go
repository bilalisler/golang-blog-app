package database

import (
	"fmt"
	model "github.com/go-blog-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:test@tcp(127.0.0.1:33077)/blog?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if DB.Migrator().HasTable("users") == false {
		DB.AutoMigrate(&model.User{})
	}

	if DB.Migrator().HasTable("posts") == false {
		DB.AutoMigrate(&model.Post{})
	}

	if DB.Migrator().HasTable("comments") == false {
		DB.AutoMigrate(&model.Comment{})
	}

	if DB.Migrator().HasTable("categories") == false {
		DB.AutoMigrate(&model.Category{})
	}

	fmt.Println("Successfully connected!")
}
