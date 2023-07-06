package main

import (
	"go-crud/controllers"
	"go-crud/initializers"
	"go-crud/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Use(middleware.RequireAuth)
	r.POST("/news", controllers.NewsCreate)
	r.GET("/newsData", controllers.NewsIndex)
	r.GET("/newsShow/:id", controllers.NewsShow)
	r.PUT("/newsUpdate/:id", controllers.NewsUpdate)
	r.DELETE("/newsDelete/:id", controllers.NewsDelete)

	r.Run()
}
