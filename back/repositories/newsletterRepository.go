package repositories

import "newsletter-back/entities"

type NewsletterRepository interface {
	GetByID(id string) (*entities.Newsletter, error)
	GetAll() ([]*entities.Newsletter, error)
	Save(newsletter *entities.Newsletter) error
	Update(newsletter *entities.Newsletter) error
	Delete(id string) error
}
