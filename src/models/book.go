package models

import "gorm.io/gorm"

// Book represent a book entry
type Book struct {
	gorm.Model
	UserID uint   `gorm:"not null;index"`
	Title  string `gorm:"not null"`
	Author string `gorm:"not null"`
	Tags   string
}
