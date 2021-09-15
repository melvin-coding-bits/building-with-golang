package db

import (
	"github.com/melvinodsa/build-with-golang/user-service/config"
	"github.com/melvinodsa/build-with-golang/user-service/models"
	"gorm.io/gorm"
)

func GetUserDetails(ctx *config.AppContext, userId uint) (*models.User, error) {
	user := &models.User{Model: gorm.Model{ID: userId}}
	err := ctx.DB.First(user).Error
	return user, err
}

func CreateUser(ctx *config.AppContext, user *models.User) error {
	return ctx.DB.Create(user).Error
}
