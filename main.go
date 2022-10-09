package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-blog-app/database"
	"github.com/go-blog-app/routes"
	_ "net/http"
)

func main() {
	database.InitDB()

	router := gin.Default()
	routes.PostRoutes(router)

	router.Run(":8080")
}
