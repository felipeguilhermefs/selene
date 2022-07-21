package postgres

import "gorm.io/gorm"

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&Book{}, &User{})
}
