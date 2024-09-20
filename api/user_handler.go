package api

import (
	"github.com/de4et/command-constructor/db"
	"github.com/de4et/command-constructor/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (u *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := u.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (u *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := u.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (u *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := u.userStore.DeleteUserByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON("success")

}

func (u *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if errs := params.Validate(); len(errs) != 0 {
		return c.JSON(errs)
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	insertedUser, err := u.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)

}
