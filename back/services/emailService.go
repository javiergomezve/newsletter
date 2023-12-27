package services

type Attachment struct {
	Filename string
	Location string
}

type EmailService interface {
	SendEmail(to []string, subject string, body string, attachments []Attachment) error
}
