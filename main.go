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

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.POST("/events", controllers.EventsCreate)
	r.PUT("/events/:id", controllers.EventsUpdate)
	r.GET("/events", controllers.EventsIndex)
	r.GET("/events/:id", controllers.EventsShow)
	r.DELETE("/events/:id", controllers.EventsDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
