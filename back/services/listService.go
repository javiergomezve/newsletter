package services

import (
	"newsletter-back/entities"
	"newsletter-back/repositories"
)

type ListService struct {
	listRepository repositories.ListRepository
}

func NewListService(listRepository repositories.ListRepository) *ListService {
	return &ListService{
		listRepository,
	}
}

func (s *ListService) GetAllLists() ([]*entities.List, error) {
	return s.listRepository.GetAll()
}

func (s *ListService) GetList(id string) (*entities.List, error) {
	return s.listRepository.GetByID(id)
}

func (s *ListService) CreateList(list *entities.List) error {
	return s.listRepository.Save(list)
}

func (s *ListService) UpdateList(list *entities.List) error {
	return s.listRepository.Update(list)
}

func (s *ListService) DeleteList(id string) error {
	return s.listRepository.Delete(id)
}
