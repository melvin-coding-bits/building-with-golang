package dto

import (
	"github.com/melvinodsa/build-with-golang/user-service/models"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u User) GetModel() *models.User {
	return &models.User{
		Model: gorm.Model{ID: u.ID},
		Name:  u.Name,
		Email: u.Email,
	}
}

func FromUserModel(user *models.User) *User {
	return &User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
