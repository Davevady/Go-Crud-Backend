package initializers

import "go-crud/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.Event{})
	DB.AutoMigrate(&models.User{})
}