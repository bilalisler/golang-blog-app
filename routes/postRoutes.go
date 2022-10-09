package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/go-blog-app/controllers"
)

func PostRoutes(r *gin.Engine) {
	postRouter := r.Group("/post")
	{
		postRouter.GET("/list", controller.AllPosts())
		postRouter.POST("/create", controller.CreatePost())
		postRouter.GET("/:id", controller.GetPost())
		postRouter.DELETE("/remove/:id", controller.RemovePost())
		postRouter.PUT("/update/:id", controller.UpdatePost())
	}
}
