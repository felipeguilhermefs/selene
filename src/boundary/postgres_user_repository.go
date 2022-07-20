package boundary

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"not null;unique_index"`
	Password string `gorm:"-"`
	Secret   string `gorm:"not null"`
}
