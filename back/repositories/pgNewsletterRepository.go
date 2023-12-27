package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"newsletter-back/entities"
	"newsletter-back/models"
)

type PgNewsletterRepository struct {
	db *gorm.DB
}

func NewPgNewsletterRepository(db *gorm.DB) *PgNewsletterRepository {
	return &PgNewsletterRepository{
		db,
	}
}

func (r *PgNewsletterRepository) GetByID(id string) (*entities.Newsletter, error) {
	var modelNewsletter models.Newsletter
	err := r.db.Preload("Recipients").Preload("Attachments").Where("id = ?", id).First(&modelNewsletter).Error
	if err != nil {
		return nil, err
	}

	entityNewsletter := convertModelToEntity(modelNewsletter)
	for _, recipient := range modelNewsletter.Recipients {
		entityNewsletter.Recipients = append(entityNewsletter.Recipients, entities.Recipient{
			ID:       fmt.Sprintf("%d", recipient.ID),
			FullName: recipient.FullName,
			Email:    recipient.Email,
			Status:   recipient.Status,
		})
	}

	for _, attachment := range modelNewsletter.Attachments {
		entityNewsletter.Attachments = append(entityNewsletter.Attachments, entities.Media{
			ID:          fmt.Sprintf("%d", attachment.ID),
			FileName:    attachment.FileName,
			ContentType: attachment.ContentType,
			Location:    attachment.Location,
		})
	}

	return entityNewsletter, nil
}

func (r *PgNewsletterRepository) GetAll() ([]*entities.Newsletter, error) {
	var modelNewsletters []models.Newsletter
	err := r.db.Preload("Recipients").Preload("Attachments").Find(&modelNewsletters).Error
	if err != nil {
		return nil, err
	}

	entityNewsletters := make([]*entities.Newsletter, len(modelNewsletters))
	for index, modelNewsletter := range modelNewsletters {
		entityNewsletters[index] = convertModelToEntity(modelNewsletter)

		for _, recipient := range modelNewsletter.Recipients {
			entityNewsletters[index].Recipients = append(entityNewsletters[index].Recipients, entities.Recipient{
				ID:       fmt.Sprintf("%d", recipient.ID),
				FullName: recipient.FullName,
				Email:    recipient.Email,
				Status:   recipient.Status,
			})
		}

		for _, attachment := range modelNewsletter.Attachments {
			entityNewsletters[index].Attachments = append(entityNewsletters[index].Attachments, entities.Media{
				ID:          fmt.Sprintf("%d", attachment.ID),
				FileName:    attachment.FileName,
				ContentType: attachment.ContentType,
				Location:    attachment.Location,
			})
		}
	}
	return entityNewsletters, nil
}

func (r *PgNewsletterRepository) Save(newsletter *entities.Newsletter) error {
	modelNewsletter := convertEntityToModel(newsletter)
	err := r.db.Create(modelNewsletter).Error
	if err != nil {
		return err
	}

	newsletter.ID = fmt.Sprintf("%d", modelNewsletter.ID)
	return nil
}

func (r *PgNewsletterRepository) Update(newsletter *entities.Newsletter) error {
	modelNewsletter := convertEntityToModel(newsletter)
	return r.db.Save(modelNewsletter).Error
}

func (r *PgNewsletterRepository) Delete(id string) error {
	return r.db.Delete(&models.Newsletter{}, id).Error
}

func convertModelToEntity(modelNewsletter models.Newsletter) *entities.Newsletter {
	return &entities.Newsletter{
		ID:      fmt.Sprintf("%d", modelNewsletter.ID),
		Subject: modelNewsletter.Subject,
		Content: modelNewsletter.Content,
		SendAt:  modelNewsletter.SendAt,
	}
}

func convertEntityToModel(entityNewsletter *entities.Newsletter) *models.Newsletter {
	recipients := make([]models.Recipient, len(entityNewsletter.Recipients))
	for index, recipient := range entityNewsletter.Recipients {
		recipients[index] = models.Recipient{
			ID: convertIDToUint(recipient.ID),
		}
	}

	attachments := make([]models.Media, len(entityNewsletter.Attachments))
	for index, media := range entityNewsletter.Attachments {
		attachments[index] = models.Media{
			ID: convertIDToUint(media.ID),
		}
	}

	return &models.Newsletter{
		ID:          convertIDToUint(entityNewsletter.ID),
		Subject:     entityNewsletter.Subject,
		Content:     entityNewsletter.Content,
		SendAt:      entityNewsletter.SendAt,
		Recipients:  recipients,
		Attachments: attachments,
	}
}

func convertIDToUint(id string) uint {
	var result uint
	fmt.Sscanf(id, "%d", &result)
	return result
}
