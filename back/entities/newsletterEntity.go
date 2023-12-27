package entities

import "time"

type Newsletter struct {
	ID          string      `json:"id"`
	Subject     string      `json:"subject"`
	Content     string      `json:"content"`
	SendAt      time.Time   `json:"send_at"`
	Recipients  []Recipient `json:"recipients"`
	Attachments []Media     `json:"attachments"`
}
