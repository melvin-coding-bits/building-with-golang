package routes

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/melvinodsa/build-with-golang/user-service/config"
	"github.com/melvinodsa/build-with-golang/user-service/dto"
	"github.com/melvinodsa/build-with-golang/user-service/helper/db"
)

func GetUserDetails(c *fiber.Ctx) error {
	ctx, ok := c.UserContext().Value(config.AppContextKey{}).(*config.AppContext)
	if !ok {
		return c.JSON(dto.Error(errors.New("context not found"), http.StatusInternalServerError))
	}
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return c.JSON(dto.Error(err, http.StatusBadRequest))
	}

	user, err := db.GetUserDetails(ctx, uint(userId))
	if err != nil {
		ctx.Logger.Error(err)
		return c.JSON(dto.Error(errors.New("error fetching data from db"), http.StatusInternalServerError))
	}
	return c.JSON(dto.Success(dto.FromUserModel(user)))
}

func GetAllUsers(c *fiber.Ctx) error {
	ctx, ok := c.UserContext().Value(config.AppContextKey{}).(*config.AppContext)
	if !ok {
		return c.JSON(dto.Error(errors.New("context not found"), http.StatusInternalServerError))
	}

	users, err := db.GetAllUsers(ctx)
	if err != nil {
		ctx.Logger.Error(err)
		return c.JSON(dto.Error(errors.New("error fetching data from db"), http.StatusInternalServerError))
	}
	return c.JSON(dto.Success(dto.FromUserModels(users)))
}

func CreateUser(c *fiber.Ctx) error {
	ctx, ok := c.UserContext().Value(config.AppContextKey{}).(*config.AppContext)
	if !ok {
		return c.JSON(dto.Error(errors.New("context not found"), http.StatusInternalServerError))
	}
	u := &dto.User{}
	if err := c.BodyParser(u); err != nil {
		return c.JSON(dto.Error(err, http.StatusBadRequest))
	}

	user := u.GetModel()
	err := db.CreateUser(ctx, user)
	if err != nil {
		return c.JSON(dto.Error(err, http.StatusInternalServerError))
	}
	return c.JSON(dto.Success(dto.FromUserModel(user)))
}

func UpdateUser(c *fiber.Ctx) error {
	ctx, ok := c.UserContext().Value(config.AppContextKey{}).(*config.AppContext)
	if !ok {
		return c.JSON(dto.Error(errors.New("context not found"), http.StatusInternalServerError))
	}
	u := &dto.User{}
	if err := c.BodyParser(u); err != nil {
		return c.JSON(dto.Error(err, http.StatusBadRequest))
	}

	user := u.GetModel()
	err := db.UpdateUser(ctx, user)
	if err != nil {
		return c.JSON(dto.Error(err, http.StatusInternalServerError))
	}
	return c.JSON(dto.Success(dto.FromUserModel(user)))
}
