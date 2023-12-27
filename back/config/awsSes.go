package config

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"newsletter-back/services"
)

type SESService struct {
	service *ses.SES
	from    string
}

func NewSESService(region, from string) (*SESService, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	service := ses.New(sess)

	return &SESService{service, from}, nil
}

func (s *SESService) SendEmail(to []string, subject string, body string, attachments []services.Attachment) error {
	message := s.createMessage(to, subject, body, attachments)

	_, err := s.service.SendRawEmail(message)
	if err != nil {
		return err
	}

	return nil
}

func (s *SESService) createMessage(to []string, subject string, body string, attachments []services.Attachment) *ses.SendRawEmailInput {
	messageData := fmt.Sprintf("From: %s\nTo: %s\r\nSubject: %s\r\n\r\n%s", s.from, to[0], subject, body)

	for _, attachment := range attachments {
		messageData += s.createAttachment(attachment)
	}

	message := &ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{
			Data: []byte(messageData),
		},
	}

	return message
}

func (s *SESService) createAttachment(attachment services.Attachment) string {
	content, err := s.downloadFile(attachment.Location)
	if err != nil {
		return ""
	}

	c := base64.StdEncoding.EncodeToString(content)
	return fmt.Sprintf("\r\nContent-Type: application/octet-stream\r\nContent-Disposition: attachment; filename=\"%s\"\r\nContent-Transfer-Encoding: base64\r\n\r\n%s",
		attachment.Filename, c)
}

func (s *SESService) downloadFile(url string) ([]byte, error) {
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
