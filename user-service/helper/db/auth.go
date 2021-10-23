package db

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/melvinodsa/build-with-golang/user-service/config"
	"github.com/melvinodsa/build-with-golang/user-service/models"
)

func CheckPassword(ctx *config.AppContext, email, password string) (*models.User, error) {
	hashedPass := sha256.Sum256([]byte(password))
	user := &models.User{}
	err := ctx.DB.First(user, "email = ? and password = ?", email, hex.EncodeToString(hashedPass[:])).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
