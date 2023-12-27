package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"newsletter-back/entities"
	"newsletter-back/models"
)

type PostgresAuthRepository struct {
	db *gorm.DB
}

func NewPostgresAuthRepository(db *gorm.DB) *PostgresAuthRepository {
	return &PostgresAuthRepository{
		db: db,
	}
}

func (repo *PostgresAuthRepository) GetUserByEmail(email string) (*entities.User, error) {
	userEntity := entities.User{}

	var userModel models.User
	r := repo.db.Where("email = ?", email).First(&userModel)
	if r.RowsAffected == 0 {
		return nil, ErrRecordNotFound
	}
	if r.Error != nil {
		return &userEntity, r.Error
	}

	userEntity.ID = fmt.Sprintf("%d", userModel.ID)
	userEntity.Name = userModel.Name
	userEntity.Email = userModel.Email
	userEntity.Password = userModel.Password

	return &userEntity, nil
}
