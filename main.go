package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	// r.GET("/", controllers.PostsCreate)

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.POST("/news", controllers.NewsCreate)
	r.GET("/newsData", controllers.NewsIndex)
	r.GET("/newsShow/:id", controllers.NewsShow)
	r.PUT("/newsUpdate/:id", controllers.NewsUpdate)
	r.DELETE("/newsDelete/:id", controllers.NewsDelete)

	r.Run()
}
