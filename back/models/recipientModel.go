package models

import "gorm.io/gorm"

type Recipient struct {
	ID          uint `gorm:"primary_key"`
	FullName    string
	Email       string       `gorm:"not null, unique"`
	Status      string       `gorm:"not null, unique"`
	Newsletters []Newsletter `gorm:"many2many:newsletter_recipient;"`
	gorm.Model
}
