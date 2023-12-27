package repositories

import "newsletter-back/entities"

type MediaRepository interface {
	GetByID(id string) (*entities.Media, error)
	GetAll() ([]*entities.Media, error)
	Save(media *entities.Media) error
	Update(media *entities.Media) error
	Delete(id string) error
}
