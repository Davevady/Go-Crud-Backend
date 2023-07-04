package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"time"

	"github.com/gin-gonic/gin"
)

func NewsCreate(c *gin.Context) {
	// Get data
	var body struct {
		Title     string
		Content   string
		Author    string
		CreatedBy string
	}

	c.Bind(&body)

	// Create a news
	now := time.Now()
	news := models.News{
		Title:     body.Title,
		Content:   body.Content,
		Author:    body.Author,
		CreatedBy: body.CreatedBy,
		CreatedAt: now,
		Date:      now,
	}

	result := initializers.DB.Create(&news)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return
	c.JSON(200, gin.H{
		"news": news,
	})
}

func NewsIndex(c *gin.Context) {
	// Get news
	var news []models.News
	initializers.DB.Find(&news)

	// Response
	c.JSON(200, gin.H{
		"news": news,
	})
}

func NewsShow(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get the news
	var news models.News
	initializers.DB.First(&news, id)

	// Response
	c.JSON(200, gin.H{
		"news": news,
	})
}

func NewsUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data
	var body struct {
		Title     string
		Content   string
		Author    string
		CreatedBy string
	}

	c.Bind(&body)

	// find
	var news models.News
	initializers.DB.First(&news, id)

	//Update
	initializers.DB.Model(&news).Updates(models.News{
		Title:     body.Title,
		Content:   body.Content,
		Author:    body.Author,
		CreatedBy: body.CreatedBy,
	})

	// Response
	c.JSON(200, gin.H{
		"news": news,
	})
}

func NewsDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	initializers.DB.Delete(&models.News{}, id)

	// Response
	c.Status(200)
}
