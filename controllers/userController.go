package controllers

import (
	"github.com/gin-gonic/gin"
	repository "github.com/go-blog-app/repositories"
	"net/http"
)

func UserPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := repository.UserPosts(1)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}

		c.JSON(http.StatusOK, posts)
	}
}
