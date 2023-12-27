package models

import "gorm.io/gorm"

type List struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Status      string `gorm:"not null"`
	Description string
	gorm.Model
}
