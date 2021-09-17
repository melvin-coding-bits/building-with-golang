package db

import (
	"github.com/melvinodsa/build-with-golang/user-service/config"
	"github.com/melvinodsa/build-with-golang/user-service/models"
	"gorm.io/gorm"
)

func GetAllUsers(ctx *config.AppContext) ([]models.User, error) {
	var users []models.User
	err := ctx.DB.Find(&users).Error
	return users, err
}

func GetUserDetails(ctx *config.AppContext, userId uint) (*models.User, error) {
	user := &models.User{Model: gorm.Model{ID: userId}}
	err := ctx.DB.First(user).Error
	return user, err
}

func CreateUser(ctx *config.AppContext, user *models.User) error {
	return ctx.DB.Create(user).Error
}

func UpdateUser(ctx *config.AppContext, user *models.User) error {
	return ctx.DB.Model(user).UpdateColumns(map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	}).Error
}
