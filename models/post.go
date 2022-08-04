package models

import (
	"github.com/go-blog-app/database"
	"log"
)

type Post struct {
	Id       int
	Body     string
	username string
}

func UserPosts(username string) []Post {
	rows, err := database.DB.Query("select id,body,username from posts where username = ?", username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var (
			post     Post
			id       int
			body     string
			username string
		)
		err = rows.Scan(&id, &body, &username)
		if err != nil {
			panic(err)
		}
		post = Post{
			id,
			body,
			username,
		}

		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return posts
}
