package services

import (
	"errors"
	"log"

	"newsletter-back/entities"
	"newsletter-back/repositories"
)

type MediaService struct {
	mediaRepository repositories.MediaRepository
	mediaUploader   MediaUploader
}

func NewMediaService(mediaRepository repositories.MediaRepository, mediaUploader MediaUploader) *MediaService {
	return &MediaService{
		mediaRepository,
		mediaUploader,
	}
}

func (s *MediaService) GetAllMedia() ([]*entities.Media, error) {
	return s.mediaRepository.GetAll()
}

func (s *MediaService) GetMediaByID(id string) (*entities.Media, error) {
	media, err := s.mediaRepository.GetByID(id)
	if errors.Is(repositories.ErrRecordNotFound, err) {
		log.Println("media does not exists")
		return nil, ErrRecordNotFound
	}
	if err != nil {
		log.Println("error getting media: ", err)
		return nil, err
	}

	return media, nil
}

func (s *MediaService) CreateMedia(media *entities.Media) error {
	fileUrl, err := s.mediaUploader.UploadFile(media.FileName, media.Content)
	if err != nil {
		log.Println("error uploading file: ", err)
		return err
	}

	media.Location = fileUrl

	return s.mediaRepository.Save(media)
}

func (s *MediaService) UpdateMedia(media *entities.Media) error {
	oldMedia, err := s.mediaRepository.GetByID(media.ID)
	if errors.Is(repositories.ErrRecordNotFound, err) {
		log.Println("media does not exists")
		return ErrRecordNotFound
	}
	if err != nil {
		log.Println("error getting media: ", err)
		return err
	}

	err = s.mediaUploader.DeleteFile(oldMedia.Location)
	if err != nil {
		log.Println("error deleting media: ", err)
		return err
	}

	fileUrl, err := s.mediaUploader.UploadFile(media.FileName, media.Content)
	if err != nil {
		log.Println("error uploading file: ", err)
		return err
	}

	media.Location = fileUrl

	return s.mediaRepository.Update(media)
}

func (s *MediaService) DeleteMediaByID(id string) error {
	media, err := s.mediaRepository.GetByID(id)
	if errors.Is(repositories.ErrRecordNotFound, err) {
		log.Println("media does not exists")
		return ErrRecordNotFound
	}
	if err != nil {
		log.Println("error getting media: ", err)
		return err
	}

	err = s.mediaUploader.DeleteFile(media.Location)
	if err != nil {
		log.Println("error deleting media: ", err)
		return err
	}

	return s.mediaRepository.Delete(id)
}
