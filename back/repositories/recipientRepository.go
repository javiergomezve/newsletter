package repositories

import "newsletter-back/entities"

type RecipientRepository interface {
	GetByID(id string) (*entities.Recipient, error)
	GetByEmail(email string) (*entities.Recipient, error)
	GetAll() ([]*entities.Recipient, error)
	Save(recipients []*entities.Recipient) error
	Update(recipient *entities.Recipient) error
	Delete(id string) error
	GetSubscribers() ([]*entities.Recipient, error)
	Unsubscribe(email string) error
}
