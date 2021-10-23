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

func Login(c *fiber.Ctx) error {
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
