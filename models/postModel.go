package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Gender   string
	Address  string
	Tempat   string
	TLahir   string
}
