package auth

import (
	"github.com/allanmarzuki/web_bkpsdm.git/internal/database"
	"github.com/allanmarzuki/web_bkpsdm.git/internal/models"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}