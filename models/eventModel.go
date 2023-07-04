package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name string
	Content string
	PostBy string
	Category string
	Date *time.Time
}