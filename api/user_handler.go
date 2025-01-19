package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/de4et/command-constructor/db"
	"github.com/de4et/command-constructor/types"
	"github.com/gofiber/fiber/v2"
)

const tokenCookieName = "apiToken"

type UserHandler struct {
	Store *db.Store
}

func NewUserHandler(store *db.Store) *UserHandler {
	return &UserHandler{
		Store: store,
	}
}

type AuthParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AuthResponseParams struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

func (u *UserHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var params AuthParams
	if err := c.BodyParser(&params); err != nil {
		return ErrWrongParams()
	}

	user, err := u.Store.User.GetUserByName(c.Context(), params.Name)
	if err != nil {
		return ErrInvalidCredentials()
	}
	if err := user.IsValidPassword(params.Password); err != nil {
		return ErrInvalidCredentials()
	}

	resp := AuthResponseParams{
		User:  user,
		Token: makeTokenFromUser(user),
	}

	// set apiToken to cookie
	c.Cookie(prepareApiTokenCookie(resp.Token))

	return c.JSON(resp)
}

func (u *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	user := c.Context().Value("user").(*types.User)
	if user == nil {
		return fmt.Errorf("no user")
	}

	return c.JSON(user)
}

func (u *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	user := c.Context().Value("user").(*types.User)
	if user == nil {
		return fmt.Errorf("no user")
	}

	err := u.Store.User.DeleteUserByID(c.Context(), user.ID)
	if err != nil {
		return err
	}
	return c.JSON("success")
}

func (u *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return ErrInvalidData()
	}

	if errs := params.Validate(); len(errs) != 0 {
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return ErrPrivateInternal()
	}

	ok, err := u.Store.User.IsExist(c.Context(), user.Name)
	if err != nil {
		return ErrPrivateInternal()
	}

	if ok {
		return ErrAlreadyExists()
	}

	insertedUser, err := u.Store.User.InsertUser(c.Context(), user)
	if err != nil {
		return ErrPrivateInternal()
	}

	c.Cookie(prepareApiTokenCookie(makeTokenFromUser(user)))

	return c.JSON(insertedUser)
}

func prepareApiTokenCookie(token string) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = tokenCookieName
	cookie.Value = token
	cookie.Expires = time.Now().AddDate(1, 0, 0)
	return cookie
}
