package models

import (
	"gorm.io/gorm"
)

// User struct to represent user model
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

// Create method to create a new user in database
func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

// Read method to read a user by id from database
func (u *User) Read(db *gorm.DB) error {
	return db.First(u, u.ID).Error
}

// Update method to update a user in database
func (u *User) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

// Delete method to delete a user from database
func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}
