package repositories

import "newsletter-back/entities"

type ListRepository interface {
	GetByID(id string) (*entities.List, error)
	GetAll() ([]*entities.List, error)
	Save(list *entities.List) error
	Update(list *entities.List) error
	Delete(id string) error
}
