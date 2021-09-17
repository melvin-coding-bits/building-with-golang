package dto

import (
	"github.com/melvinodsa/build-with-golang/user-service/models"
	"gorm.io/gorm"
)

type User struct {
	ID    uint        `json:"id"`
	Role  models.Role `json:"role"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
}

func (u User) GetModel() *models.User {
	return &models.User{
		Model: gorm.Model{ID: u.ID},
		Role:  u.Role,
		Name:  u.Name,
		Email: u.Email,
	}
}

func FromUserModel(user *models.User) *User {
	return &User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}

func FromUserModels(users []models.User) []User {
	result := make([]User, len(users))
	for i, user := range users {
		result[i] = *(FromUserModel(&user))
	}
	return result
}
