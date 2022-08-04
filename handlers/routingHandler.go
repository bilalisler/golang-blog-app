package handlers

import (
	"fmt"
	model "github.com/go-blog-app/models"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var posts []model.Post
	posts = model.UserPosts("test2")

	fmt.Println(posts)

}
