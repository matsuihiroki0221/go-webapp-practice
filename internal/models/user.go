package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
