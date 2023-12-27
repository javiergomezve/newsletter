package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"newsletter-back/entities"
	"newsletter-back/models"
)

type PgSubscriberRepository struct {
	db *gorm.DB
}

func NewPgRecipientRepository(db *gorm.DB) *PgSubscriberRepository {
	return &PgSubscriberRepository{
		db: db,
	}
}

func (r *PgSubscriberRepository) GetAll() ([]*entities.Recipient, error) {
	var modelSubscribers []models.Recipient
	err := r.db.Find(&modelSubscribers).Error
	if err != nil {
		return nil, err
	}

	entitySubscribers := make([]*entities.Recipient, len(modelSubscribers))
	for index, modelSubscriber := range modelSubscribers {
		entitySubscribers[index] = convertModelToEntityRecipient(modelSubscriber)
	}
	return entitySubscribers, nil
}

func (r *PgSubscriberRepository) GetByID(id string) (*entities.Recipient, error) {
	var modelSubscriber models.Recipient
	if err := r.db.Where("id = ?", id).First(&modelSubscriber).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Suscriptor no encontrado")
		}
		return nil, err
	}

	entitySubscriber := convertModelToEntityRecipient(modelSubscriber)
	return entitySubscriber, nil
}

func (r *PgSubscriberRepository) GetByEmail(email string) (*entities.Recipient, error) {
	var modelSubscriber models.Recipient
	err := r.db.Where("email = ?", email).First(&modelSubscriber).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	entitySubscriber := convertModelToEntityRecipient(modelSubscriber)
	return entitySubscriber, nil
}

func (r *PgSubscriberRepository) Save(recipients []*entities.Recipient) error {
	mRecipients := make([]models.Recipient, len(recipients))

	for index, recipient := range recipients {
		mRecipients[index] = models.Recipient{
			FullName: recipient.FullName,
			Email:    recipient.Email,
			Status:   recipient.Status,
		}
	}

	err := r.db.Create(mRecipients).Error
	if err != nil {
		return err
	}

	for index, recipient := range mRecipients {
		recipients[index].ID = fmt.Sprintf("%d", recipient.ID)
	}

	return nil
}

func (r *PgSubscriberRepository) Update(subscriber *entities.Recipient) error {
	modelSubscriber := convertEntityToModelRecipient(subscriber)
	return r.db.Save(modelSubscriber).Error
}

func (r *PgSubscriberRepository) Delete(id string) error {
	return r.db.Delete(&entities.Recipient{}, id).Error
}

func convertModelToEntityRecipient(modelRecipient models.Recipient) *entities.Recipient {
	return &entities.Recipient{
		ID:       fmt.Sprintf("%d", modelRecipient.ID),
		FullName: modelRecipient.FullName,
		Email:    modelRecipient.Email,
		Status:   modelRecipient.Status,
	}
}

func convertEntityToModelRecipient(entityRecipient *entities.Recipient) *models.Recipient {
	return &models.Recipient{
		ID:       convertIDToUint(entityRecipient.ID),
		FullName: entityRecipient.FullName,
		Email:    entityRecipient.Email,
		Status:   entityRecipient.Status,
	}
}
