//Package db has the helper function to interact with the database
package db

import (
	"github.com/melvinodsa/build-with-golang/user-service/config"
	"github.com/melvinodsa/build-with-golang/user-service/models"
	"gorm.io/gorm"
)

//GetAllUsers returns all the users in db. This is intended for admin
func GetAllUsers(ctx *config.AppContext) ([]models.User, error) {
	var users []models.User
	err := ctx.DB.Find(&users).Error
	return users, err
}

//GetUserDetails returns the user details for a given id
func GetUserDetails(ctx *config.AppContext, userId uint) (*models.User, error) {
	user := &models.User{Model: gorm.Model{ID: userId}}
	err := ctx.DB.First(user).Error
	return user, err
}

//CreateUser creates a new user in db
func CreateUser(ctx *config.AppContext, user *models.User) error {
	return ctx.DB.Create(user).Error
}

//UpdateUser updates the user mode in db
func UpdateUser(ctx *config.AppContext, user *models.User) error {
	return ctx.DB.Model(user).UpdateColumns(map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	}).Error
}
