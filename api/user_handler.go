package api

import (
	"github.com/de4et/command-constructor/db"
	"github.com/de4et/command-constructor/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Store *db.Store
}

func NewUserHandler(store *db.Store) *UserHandler {
	return &UserHandler{
		Store: store,
	}
}

type AuthParams struct {
	Name     string
	Password string
}

type AuthResponseParams struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

func (u *UserHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var params AuthParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	user, err := u.Store.User.GetUserByName(c.Context(), params.Name)
	if err != nil {
		return err
	}
	if err := user.IsValidPassword(params.Password); err != nil {
		return err
	}

	resp := AuthResponseParams{
		User:  user,
		Token: makeTokenFromUser(user),
	}
	return c.JSON(resp)
}

func (u *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := u.Store.User.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (u *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := u.Store.User.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (u *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := u.Store.User.DeleteUserByID(c.Context(), id)
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

	insertedUser, err := u.Store.User.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)

}
