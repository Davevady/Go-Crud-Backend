package models

import "time"

type News struct {
	ID        int       `gorm:"primaryKey"`
	Title     string    `gorm:"size:255"`
	Content   string    `gorm:"type:text"`
	Author    string    `gorm:"size:255"`
	Date      time.Time `gorm:"type:date"`
	CreatedBy string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"type:timestamp"`
}
