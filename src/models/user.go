package models

import "gorm.io/gorm"

// User represent an user of the system
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"not null;unique_index"`
	Password string `gorm:"-"`
	Secret   string `gorm:"not null"`
}
