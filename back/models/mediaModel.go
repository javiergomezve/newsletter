package models

import "gorm.io/gorm"

type Media struct {
	ID          uint         `gorm:"primary_key"`
	FileName    string       `gorm:"not null"`
	ContentType string       `gorm:"not null"`
	Location    string       `gorm:"not null"`
	Newsletters []Newsletter `gorm:"many2many:newsletter_attachment;"`
	gorm.Model
}
