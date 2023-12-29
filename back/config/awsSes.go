package config

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"newsletter-back/services"
)

type SESService struct {
	service     *ses.SES
	from        string
	frontendUrl string
}

func NewSESService(region, from, frontendUrl string) (*SESService, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	service := ses.New(sess)

	return &SESService{service, from, frontendUrl}, nil
}

func (s *SESService) SendEmail(to []string, subject string, body string, attachments []services.Attachment) error {
	for _, recipient := range to {
		log.Println("recipient: ", recipient)

		message := s.createMessage(recipient, subject, body, attachments)
		_, err := s.service.SendRawEmail(message)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SESService) createMessage(to string, subject string, body string, attachments []services.Attachment) *ses.SendRawEmailInput {
	unsubscribeLink := fmt.Sprintf("<a href=\"%s/unsubscribe?email=%s\">Unsubscribe</a>", s.frontendUrl, to)

	boundary := "MyBoundary"
	message := fmt.Sprintf(
		"From: %s\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: multipart/mixed; boundary=\"%s\"\r\n\r\n",
		s.from, to, subject, boundary,
	)

	message += fmt.Sprintf("--%s\r\n", boundary)
	message += "Content-Type: text/html; charset=utf-8\r\n\r\n"
	message += body + unsubscribeLink
	message += "\r\n"

	for _, attachment := range attachments {
		message += fmt.Sprintf("--%s\r\n", boundary)
		message += s.createAttachment(attachment)
		message += "\r\n"
	}

	message += fmt.Sprintf("--%s--\r\n", boundary)

	rawMessage := &ses.RawMessage{Data: []byte(message)}

	return &ses.SendRawEmailInput{RawMessage: rawMessage}
}

func (s *SESService) createAttachment(attachment services.Attachment) string {
	content, err := s.downloadFile(attachment.Location)
	if err != nil {
		return ""
	}

	c := base64.StdEncoding.EncodeToString(content)
	result := fmt.Sprintf("Content-Type: application/octet-stream\r\n")
	result += fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", attachment.Filename)
	result += "Content-Transfer-Encoding: base64\r\n\r\n"
	result += c

	return result
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
