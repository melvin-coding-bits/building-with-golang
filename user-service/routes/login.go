package routes

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/melvinodsa/build-with-golang/user-service/config"
	"github.com/melvinodsa/build-with-golang/user-service/crypto_utils"
	"github.com/melvinodsa/build-with-golang/user-service/dto"
	"github.com/melvinodsa/build-with-golang/user-service/helper/db"
)

//Login logs the user into the system
func Login(c *fiber.Ctx) error {
	/*
	 * We will get the user context
	 * We will get the auth payload to be checked
	 * We will validate the user email and password
	 * We will generate the token and set as auth token
	 */
	ctx, ok := c.UserContext().Value(config.AppContextKey{}).(*config.AppContext)
	if !ok {
		return c.JSON(dto.Error(errors.New("context not found"), http.StatusInternalServerError))
	}
	a := &dto.Auth{}
	if err := c.BodyParser(a); err != nil {
		return c.JSON(dto.Error(err, http.StatusBadRequest))
	}

	user, err := db.CheckPassword(ctx, a.Email, a.Password)
	if err != nil {
		return c.JSON(dto.Error(err, http.StatusForbidden))
	}

	token, err := crypto_utils.GenerateToken(32)
	if err != nil {
		return c.JSON(dto.Error(err, http.StatusInternalServerError))
	}
	c.Response().Header.Set("Authorization", "Bearer "+token)

	return c.JSON(dto.Success(dto.FromUserModel(user)))
}
