package models

import "gorm.io/gorm"

type BlogPost struct {
	gorm.Model
	Title   string
	Content string
}
