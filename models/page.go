package models

import "gorm.io/gorm"

type Page struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
}
