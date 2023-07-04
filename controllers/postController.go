package controllers

import (
	"tugas1/initializers"
	"tugas1/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get dataoff req body
	var body struct {
		Name    string
		Email   string
		Gender  string
		Address string
		Tempat  string
		TLahir  string
	}

	c.Bind((&body))

	// Create a post
	post := models.Post{Name: body.Name, Email: body.Email, Gender: body.Gender,
		Address: body.Address, Tempat: body.Tempat, TLahir: body.TLahir}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Name    string
		Email   string
		Gender  string
		Address string
		Tempat  string
		TLahir  string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Updated it
	initializers.DB.Model(&post).Updates(models.Post{
		Name:    body.Name,
		Email:   body.Email,
		Gender:  body.Gender,
		Address: body.Address,
		Tempat:  body.Tempat,
		TLahir:  body.TLahir,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
