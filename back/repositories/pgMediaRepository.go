package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"newsletter-back/entities"
	"newsletter-back/models"
)

type PgMediaRepository struct {
	db *gorm.DB
}

func NewPgMediaRepository(db *gorm.DB) *PgMediaRepository {
	return &PgMediaRepository{
		db: db,
	}
}

func (r *PgMediaRepository) GetAll() ([]*entities.Media, error) {
	var modelMedias []models.Media

	err := r.db.Find(&modelMedias).Error
	if err != nil {
		return nil, err
	}

	entityMedias := make([]*entities.Media, len(modelMedias))
	for index, modelMedia := range modelMedias {
		entityMedias[index] = convertModelToEntityMedia(modelMedia)
	}
	return entityMedias, nil
}

func (r *PgMediaRepository) GetByID(id string) (*entities.Media, error) {
	var modelMedia models.Media

	err := r.db.Where("id = ?", id).First(&modelMedia).Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}

	entityMedia := convertModelToEntityMedia(modelMedia)
	return entityMedia, nil
}

func (r *PgMediaRepository) Save(media *entities.Media) error {
	modelMedia := convertEntityToModelMedia(media)
	err := r.db.Create(modelMedia).Error
	if err != nil {
		return err
	}

	media.ID = fmt.Sprintf("%d", modelMedia.ID)
	return nil
}

func (r *PgMediaRepository) Update(media *entities.Media) error {
	modelMedia := convertEntityToModelMedia(media)
	return r.db.Save(modelMedia).Error
}

func (r *PgMediaRepository) Delete(id string) error {
	return r.db.Delete(&models.Media{}, id).Error
}

func convertModelToEntityMedia(modelMedia models.Media) *entities.Media {
	return &entities.Media{
		ID:          fmt.Sprintf("%d", modelMedia.ID),
		FileName:    modelMedia.FileName,
		ContentType: modelMedia.ContentType,
		Location:    modelMedia.Location,
	}
}

func convertEntityToModelMedia(entityMedia *entities.Media) *models.Media {
	return &models.Media{
		ID:          convertIDToUint(entityMedia.ID),
		FileName:    entityMedia.FileName,
		ContentType: entityMedia.ContentType,
		Location:    entityMedia.Location,
	}
}
