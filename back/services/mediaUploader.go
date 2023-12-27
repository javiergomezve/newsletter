package services

import "io"

type MediaUploader interface {
	UploadFile(name string, file io.Reader) (string, error)
	DeleteFile(fileURL string) error
	GetSignedURL(fileURL string) (string, error)
}
