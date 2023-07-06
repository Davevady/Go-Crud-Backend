package initializers

import "tugas1/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Post{})
}
