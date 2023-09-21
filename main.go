package main

import (
	"project-ankit/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/posts", handlers.GetBlogPosts)
	r.GET("/posts/:id", handlers.GetBlogPost)
	r.POST("/posts", handlers.CreateBlogPost)
	r.PUT("/posts/:id", handlers.UpdateBlogPost)
	r.DELETE("/posts/:id", handlers.DeleteBlogPost)
	r.Run(":8002")
}
