package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"newsletter-back/entities"
	"newsletter-back/models"
)

type PgListRepository struct {
	db *gorm.DB
}

func NewPgListRepository(db *gorm.DB) *PgListRepository {
	return &PgListRepository{
		db: db,
	}
}

func (r *PgListRepository) GetByID(id string) (*entities.List, error) {
	var modelList entities.List
	if err := r.db.Where("id = ?", id).First(&modelList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("list does not exists")
		}
		return nil, err
	}

	entityList := convertModelToEntityList(modelList)
	return entityList, nil
}

func (r *PgListRepository) GetAll() ([]*entities.List, error) {
	var modelLists []entities.List
	if err := r.db.Find(&modelLists).Error; err != nil {
		return nil, err
	}

	var entityLists []*entities.List
	for _, modelList := range modelLists {
		entityLists = append(entityLists, convertModelToEntityList(modelList))
	}
	return entityLists, nil
}

func (r *PgListRepository) Save(list *entities.List) error {
	modelList := convertEntityToModelList(list)
	return r.db.Create(modelList).Error
}

func (r *PgListRepository) Update(list *entities.List) error {
	modelList := convertEntityToModelList(list)
	return r.db.Save(modelList).Error
}

func (r *PgListRepository) Delete(id string) error {
	return r.db.Delete(&entities.List{}, id).Error
}

func convertModelToEntityList(modelList entities.List) *entities.List {
	return &entities.List{
		ID:          fmt.Sprintf("%d", modelList.ID),
		Name:        modelList.Name,
		Status:      modelList.Status,
		Description: modelList.Description,
	}
}

func convertEntityToModelList(entityList *entities.List) *models.List {
	return &models.List{
		ID:          convertIDToUint(entityList.ID),
		Name:        entityList.Name,
		Status:      entityList.Status,
		Description: entityList.Description,
	}
}
