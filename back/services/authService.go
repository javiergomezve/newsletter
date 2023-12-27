package services

import (
	"log"

	"newsletter-back/repositories"
)

type SignToken func(payload interface{}) (string, error)
type ComparePasswords func(hashedPassword string, password string) error

type AuthService struct {
	authRepository   repositories.AuthRepository
	signToken        SignToken
	comparePasswords ComparePasswords
}

func NewAuthService(authRepository repositories.AuthRepository, signToken SignToken, comparePasswords ComparePasswords) *AuthService {
	return &AuthService{
		authRepository,
		signToken,
		comparePasswords,
	}
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.authRepository.GetUserByEmail(email)
	if err != nil {
		log.Println("user does not exists")
		return "", err
	}

	err = s.comparePasswords(user.Password, password)
	if err != nil {
		log.Println("wrong password")
		return "", err
	}

	token, err := s.signToken(user)
	if err != nil {
		log.Println("could not sign token")
		return "", err
	}

	return token, nil
}
