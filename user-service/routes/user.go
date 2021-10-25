package routes

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/melvinodsa/build-with-golang/user-service/config"
	"github.com/melvinodsa/build-with-golang/user-service/dto"
	"github.com/melvinodsa/build-with-golang/user-service/helper/db"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var payloadParsingError = promauto.NewCounter(prometheus.CounterOpts{
	Name: "payload_parsing_error",
	Help: "Errors while parsing the payload",
})

//GetUserDetails returns the details of a user for a given user id
func GetUserDetails(c *fiber.Ctx) error {
	/*
	 * We will get the user context
	 * We will get the user id from the request
	 * We will get the user details from the db
	 */
	ctx, ok := c.UserContext().Value(config.AppContextKey{}).(*config.AppContext)
	if !ok {
		return c.JSON(dto.Error(errors.New("context not found"), http.StatusInternalServerError))
	}
	userId, err := c.ParamsInt("userId")
	if err != nil {
		payloadParsingError.Inc()
		return c.JSON(dto.Error(err, http.StatusBadRequest))
	}

	user, err := db.GetUserDetails(ctx, uint(userId))
	if err != nil {
		ctx.Logger.Error(err)
		return c.JSON(dto.Error(errors.New("error fetching data from db"), http.StatusInternalServerError))
	}
	return c.JSON(dto.Success(dto.FromUserModel(user)))
}

//GetAllUsers returns all the users
func GetAllUsers(c *fiber.Ctx) error {
	/*
	 * We will get the user context
	 * We will get all the users from the db
	 */
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

//CreateUser creates a new user in db
func CreateUser(c *fiber.Ctx) error {
	/*
	 * We will get the user context
	 * We will get the user payload to be created
	 * WE will validate the user payload
	 * We will create the user in the db
	 */
	ctx, ok := c.UserContext().Value(config.AppContextKey{}).(*config.AppContext)
	if !ok {
		return c.JSON(dto.Error(errors.New("context not found"), http.StatusInternalServerError))
	}

	u := &dto.User{}
	if err := c.BodyParser(u); err != nil {
		payloadParsingError.Inc()
		return c.JSON(dto.Error(err, http.StatusBadRequest))
	}

	if err := u.Validate(); err != nil {
		return c.JSON(dto.Error(err, http.StatusBadRequest))
	}

	user := u.GetModel()
	err := db.CreateUser(ctx, user)
	if err != nil {
		return c.JSON(dto.Error(err, http.StatusInternalServerError))
	}
	return c.JSON(dto.Success(dto.FromUserModel(user)))
}

//UpdateUser updates a user in db
func UpdateUser(c *fiber.Ctx) error {
	/*
	 * We will get the user context
	 * We will get the user payload to be update
	 * We will update the user in the db
	 */
	ctx, ok := c.UserContext().Value(config.AppContextKey{}).(*config.AppContext)
	if !ok {
		return c.JSON(dto.Error(errors.New("context not found"), http.StatusInternalServerError))
	}
	u := &dto.User{}
	if err := c.BodyParser(u); err != nil {
		payloadParsingError.Inc()
		return c.JSON(dto.Error(err, http.StatusBadRequest))
	}

	user := u.GetModel()
	err := db.UpdateUser(ctx, user)
	if err != nil {
		return c.JSON(dto.Error(err, http.StatusInternalServerError))
	}
	return c.JSON(dto.Success(dto.FromUserModel(user)))
}
