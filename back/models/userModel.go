package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `json:"first_name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null, unique"`
	Password string `json:"password" gorm:"not null"`
	gorm.Model
}
