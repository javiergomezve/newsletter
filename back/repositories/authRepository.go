package repositories

import "newsletter-back/entities"

type AuthRepository interface {
	GetUserByEmail(email string) (*entities.User, error)
}
