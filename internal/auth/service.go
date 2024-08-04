package auth

import (
	"errors"
	"github.com/allanmarzuki/web_bkpsdm.git/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

func (s *Service) Login(username, password string) (string, error) {
	user, err := s.Repository.FindUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Implement JWT token generation here
	return "dummy_token", nil
}

func (s *Service) Register(username, name, email, noHP, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Name:     name,
		Email:    email,
		NoHP:     noHP,
		Password: string(hashedPassword),
		RoleID:   1, // Assume default role ID is 1
	}

	return s.Repository.CreateUser(user)
}