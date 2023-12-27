package config

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"

	"newsletter-back/services"
)

type GmailService struct {
	service *gmail.Service
}

func NewGmailService(credentialsFile string) (*GmailService, error) {
	b, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		return nil, err
	}

	client := config.Client(context.Background())

	service, err := gmail.New(client)
	if err != nil {
		return nil, err
	}

	return &GmailService{service: service}, nil
}

func (g *GmailService) SendEmail(to []string, subject string, body string, attachments []services.Attachment) error {
	message := g.createMessage(to, subject, body, attachments)
	log.Println("try to send")
	_, err := g.service.Users.Messages.Send("me", message).Do()
	if err != nil {
		return err
	}
	return nil
}

func (g *GmailService) createMessage(to []string, subject string, body string, attachments []services.Attachment) *gmail.Message {
	var message gmail.Message

	var emailContent string
	emailContent += "To: " + to[0] + "\r\n"
	emailContent += "Subject: " + subject + "\r\n"
	emailContent += "\r\n" + body + "\r\n"

	message.Raw = base64.StdEncoding.EncodeToString([]byte(emailContent))

	for _, attachment := range attachments {
		message.Raw += g.createAttachment(attachment)
	}

	return &message
}

func (g *GmailService) createAttachment(attachment services.Attachment) string {
	content, err := g.downloadFile(attachment.Location)
	if err != nil {
		return ""
	}

	c := base64.StdEncoding.EncodeToString(content)
	return fmt.Sprintf("\r\nContent-Type: application/octet-stream\r\nContent-Disposition: attachment; filename=\"%s\"\r\nContent-Transfer-Encoding: base64\r\n\r\n%s",
		attachment.Filename, c)
}

func (g *GmailService) downloadFile(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	fileContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
