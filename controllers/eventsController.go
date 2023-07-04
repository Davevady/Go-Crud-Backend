package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"time"

	"github.com/gin-gonic/gin"
)

func EventsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Name     string
		Content  string
		PostBy   string
		Category string
	}
	c.Bind(&body)

	// Create a event
	now := time.Now()
	event := models.Event{
		Name:     body.Name,
		Content:  body.Content,
		PostBy:   body.PostBy,
		Category: body.Category,
		Date:     &now,
	}
	result := initializers.DB.Create(&event)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it
	c.JSON(200, gin.H{
		"event": event,
	})
}

func EventsIndex(c *gin.Context) {
	// Get the events
	var events []models.Event
	initializers.DB.Find(&events)

	// Respond with them
	c.JSON(200, gin.H{
		"event": events,
	})
}

func EventsShow(c *gin.Context) {
	// Get id off URL
	id := c.Param("id")

	// Get the events
	var event models.Event
	initializers.DB.First(&event, id)

	// Respond with them
	c.JSON(200, gin.H{
		"event": event,
	})
}

func EventsUpdate(c *gin.Context) {
	// Get the id off the URL
	id := c.Param("id")

	// Get the data off req body
	
	var body struct {
		Name     string
		Content  string
		PostBy   string
		Category string
		Date     *time.Time
	}
	c.Bind(&body)

	// Find the post were updating
	var event models.Event
	initializers.DB.First(&event, id)

	// Update it
	now := time.Now()
	initializers.DB.Model(&event).Updates(models.Event{
		Name:     body.Name,
		Content:  body.Content,
		PostBy:   body.PostBy,
		Category: body.Category,
		Date:     &now,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"event": event,
	})
}

func EventsDelete(c *gin.Context) {
	// Get the id off the URL
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Event{}, id)

	// Respond
	c.Status(200)
}
