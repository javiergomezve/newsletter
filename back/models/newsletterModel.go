package models

import (
	"time"

	"gorm.io/gorm"
)

type Newsletter struct {
	ID          uint        `gorm:"primary_key"`
	Subject     string      `gorm:"not null"`
	Content     string      `gorm:"not null"`
	SendAt      time.Time   `gorm:"not null"`
	Recipients  []Recipient `gorm:"many2many:newsletter_recipient;"`
	Attachments []Media     `gorm:"many2many:newsletter_attachment;"`
	gorm.Model
}
