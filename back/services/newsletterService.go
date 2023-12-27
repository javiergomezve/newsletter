package services

import (
	"newsletter-back/entities"
	"newsletter-back/repositories"
)

type NewsletterService struct {
	newsletterRepository repositories.NewsletterRepository
}

func NewNewsletterService(newsletterRepository repositories.NewsletterRepository) *NewsletterService {
	return &NewsletterService{
		newsletterRepository,
	}
}

func (s *NewsletterService) GetAllNewsletters() ([]*entities.Newsletter, error) {
	return s.newsletterRepository.GetAll()
}

func (s *NewsletterService) GetNewsletter(id string) (*entities.Newsletter, error) {
	return s.newsletterRepository.GetByID(id)
}

func (s *NewsletterService) CreateNewsletter(newsletter *entities.Newsletter) error {
	return s.newsletterRepository.Save(newsletter)
}

func (s *NewsletterService) UpdateNewsletter(newsletter *entities.Newsletter) error {
	return s.newsletterRepository.Update(newsletter)
}

func (s *NewsletterService) DeleteNewsletter(id string) error {
	return s.newsletterRepository.Delete(id)
}
