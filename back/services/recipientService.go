package services

import (
	"newsletter-back/entities"
	"newsletter-back/repositories"
)

type RecipientService struct {
	recipientRepository repositories.RecipientRepository
}

func NewRecipientService(recipientRepository repositories.RecipientRepository) *RecipientService {
	return &RecipientService{
		recipientRepository,
	}
}

func (s *RecipientService) GetRecipientByID(id string) (*entities.Recipient, error) {
	return s.recipientRepository.GetByID(id)
}

func (s *RecipientService) GetRecipientByEmail(email string) (*entities.Recipient, error) {
	return s.recipientRepository.GetByEmail(email)
}

func (s *RecipientService) GetAllRecipients() ([]*entities.Recipient, error) {
	return s.recipientRepository.GetAll()
}

func (s *RecipientService) CreateRecipients(recipients []*entities.Recipient) error {
	// TODO: fix
	for _, recipient := range recipients {
		recipient.Status = "subscribed"
	}
	return s.recipientRepository.Save(recipients)
}

func (s *RecipientService) UpdateRecipient(recipient *entities.Recipient) error {
	return s.recipientRepository.Update(recipient)
}

func (s *RecipientService) DeleteRecipients(id string) error {
	return s.recipientRepository.Delete(id)
}

func (s *RecipientService) GetSubscribers() ([]*entities.Recipient, error) {
	return s.recipientRepository.GetSubscribers()
}

func (s *RecipientService) Unsubscribe(email string) error {
	return s.recipientRepository.Unsubscribe(email)
}
