package services

import (
	"newsletter-back/entities"
	"newsletter-back/repositories"
)

type RecipientService struct {
	subscriberRepository repositories.RecipientRepository
}

func NewRecipientService(subscriberRepository repositories.RecipientRepository) *RecipientService {
	return &RecipientService{
		subscriberRepository,
	}
}

func (s *RecipientService) GetRecipientByID(id string) (*entities.Recipient, error) {
	return s.subscriberRepository.GetByID(id)
}

func (s *RecipientService) GetSubscriberByEmail(email string) (*entities.Recipient, error) {
	return s.subscriberRepository.GetByEmail(email)
}

func (s *RecipientService) GetAllSubscribers() ([]*entities.Recipient, error) {
	return s.subscriberRepository.GetAll()
}

func (s *RecipientService) CreateRecipients(recipients []*entities.Recipient) error {
	// TODO: fix
	for _, recipient := range recipients {
		recipient.Status = "active"
	}
	return s.subscriberRepository.Save(recipients)
}

func (s *RecipientService) UpdateSubscriber(subscriber *entities.Recipient) error {
	return s.subscriberRepository.Update(subscriber)
}

func (s *RecipientService) DeleteSubscriberByID(id string) error {
	return s.subscriberRepository.Delete(id)
}
