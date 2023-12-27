package config

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type AWSS3Uploader struct {
	Region          string
	Bucket          string
	AccessKeyID     string
	SecretAccessKey string
}

func NewAWSS3Uploader(region, bucket, accessKeyID, secretAccessKey string) *AWSS3Uploader {
	return &AWSS3Uploader{
		Region:          region,
		Bucket:          bucket,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	}
}

func (u *AWSS3Uploader) UploadFile(name string, file io.Reader) (string, error) {
	fileURL := ""

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(u.Region),
	})
	if err != nil {
		return fileURL, err
	}

	svc := s3.New(sess)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return fileURL, err
	}

	fileReader := bytes.NewReader(fileBytes)
	newFileName := generateUUIDFileName(name)

	params := &s3.PutObjectInput{
		Bucket: aws.String(u.Bucket),
		Key:    aws.String(newFileName),
		Body:   fileReader,
	}

	_, err = svc.PutObject(params)
	if err != nil {
		return fileURL, err
	}

	fileURL = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", u.Bucket, newFileName)

	return fileURL, nil
}

func (u *AWSS3Uploader) DeleteFile(fileURL string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(u.Region),
	})
	if err != nil {
		return err
	}

	svc := s3.New(sess)

	fileName := extractFileNameFromURL(fileURL)

	params := &s3.DeleteObjectInput{
		Bucket: aws.String(u.Bucket),
		Key:    aws.String(fileName),
	}

	_, err = svc.DeleteObject(params)
	if err != nil {
		return err
	}

	return nil
}

func (u *AWSS3Uploader) GetSignedURL(fileURL string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(u.Region),
	})
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	fileName := extractFileNameFromURL(fileURL)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(u.Bucket),
		Key:    aws.String(fileName),
	})
	url, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return url, nil
}

func extractFileNameFromURL(fileURL string) string {
	lastSlash := 0
	for i := len(fileURL) - 1; i >= 0; i-- {
		if fileURL[i] == '/' {
			lastSlash = i
			break
		}
	}
	return fileURL[lastSlash+1:]
}

func generateUUIDFileName(originalName string) string {
	fileExtension := filepath.Ext(originalName)

	uuidString := uuid.New().String()

	uuidFileName := "medias/" + uuidString + fileExtension

	return uuidFileName
}
