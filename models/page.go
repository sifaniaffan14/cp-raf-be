package models

import "gorm.io/gorm"

type Page struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
}
