package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Name    string
	Email   string
	Gender  string
	Address string
	Tempat  string
	TLahir  string
}
