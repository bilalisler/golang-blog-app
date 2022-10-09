package repositories

import (
	"github.com/go-blog-app/database"
	model "github.com/go-blog-app/models"
	"strconv"
)

func UserPosts(userId int) ([]model.Post, error) {
	var posts []model.Post

	result := database.DB.Model(&model.Post{}).Preload("User").Preload("Category").Find(&posts, "user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func AllPosts() ([]model.Post, error) {
	var posts []model.Post

	result := database.DB.Model(&model.Post{}).Preload("User").Preload("Category").Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func GetPost(postId string) (model.Post, error) {
	var post model.Post
	result := database.DB.Preload("User").Preload("Category").Preload("Comments.User").First(&post, postId)

	if result.Error != nil {
		return post, result.Error
	}

	return post, nil
}

func CreatePost(post model.Post) error {
	result := database.DB.Create(&post)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePost(post model.Post) error {
	_, err := GetPost(strconv.Itoa(post.ID))

	if err != nil {
		return err
	}

	result := database.DB.Save(&post)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func RemovePost(postId string) error {
	_, err := GetPost(postId)

	if err != nil {
		return err
	}

	result := database.DB.Delete(&model.Post{}, postId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
