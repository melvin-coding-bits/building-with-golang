package dto

import (
	"errors"

	"github.com/melvinodsa/build-with-golang/user-service/models"
	"gorm.io/gorm"
)

//User dto to interact with api
type User struct {
	ID       uint        `json:"id"`
	Role     models.Role `json:"role"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
}

//GetModel returns the user model from dto
func (u User) GetModel() *models.User {
	return &models.User{
		Model:    gorm.Model{ID: u.ID},
		Role:     u.Role,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u User) Validate() error {
	if len(u.Name) == 0 {
		return errors.New("name cannot be empty")
	}
	if len(u.Email) == 0 {
		return errors.New("email cannot be empty")
	}
	if len(u.Password) == 0 {
		return errors.New("password cannot be empty")
	}
	return nil
}

//FromUserModel returns the user dto from user model
func FromUserModel(user *models.User) *User {
	return &User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}

//FromUserModel returns the list of user dto from user model list
func FromUserModels(users []models.User) []User {
	result := make([]User, len(users))
	for i, user := range users {
		result[i] = *(FromUserModel(&user))
	}
	return result
}
