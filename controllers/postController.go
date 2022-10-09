package controllers

import (
	"github.com/gin-gonic/gin"
	model "github.com/go-blog-app/models"
	repository "github.com/go-blog-app/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"net/http"
	"time"
)

var validate *validator.Validate = validator.New()

func AllPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := repository.AllPosts()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}

		c.JSON(http.StatusOK, posts)
	}
}

func GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("id")
		post, err := repository.GetPost(postId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}

		c.JSON(http.StatusOK, post)
	}
}

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post model.Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		post.CreatedAt = time.Now()
		post.Slug = slug.Make(post.Title)

		err := validate.Struct(&post)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = repository.CreatePost(post)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "OK",
		})
	}
}

func RemovePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("id")

		err := repository.RemovePost(postId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}
}

func UpdatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("id")

		post, err := repository.GetPost(postId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err = c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = repository.UpdatePost(post)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"post":    post,
		})
	}
}
