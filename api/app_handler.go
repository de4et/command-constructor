package api

import (
	"github.com/a-h/templ"
	"github.com/de4et/command-constructor/db"
	"github.com/de4et/command-constructor/types"
	"github.com/de4et/command-constructor/view"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type MainHandler struct {
	Store *db.Store
}

func NewMainHandler(store *db.Store) *MainHandler {
	return &MainHandler{
		Store: store,
	}
}

func (u *MainHandler) HandleMain(c *fiber.Ctx) error {
	user, _ := c.Context().Value("user").(*types.User)
	commandTemplates := []*types.CommandTemplate{}

	if user != nil {
		var err error
		commandTemplates, err = u.Store.Command.GetCommandsOfUser(c.Context(), user.ID)
		if err != nil {
			return ErrPrivateInternal()
		}
	}

	index := view.Index(commandTemplates, user)
	handler := adaptor.HTTPHandler(templ.Handler(index))
	return handler(c)
}

func (u *MainHandler) HandleCreate(c *fiber.Ctx) error {
	user, _ := c.Context().Value("user").(*types.User)

	if user == nil {
		return c.Redirect("/main")
	}

	create := view.CreateTemplate(user, "Создание шаблона")
	handler := adaptor.HTTPHandler(templ.Handler(create))
	return handler(c)
}

func (u *MainHandler) HandleEdit(c *fiber.Ctx) error {
	user, _ := c.Context().Value("user").(*types.User)
	id := c.Params("id")

	if user == nil {
		return c.Redirect("/main")
	}

	command, err := u.Store.Command.GetCommandByID(c.Context(), id)
	if err != nil {
		return err
	}

	create := view.EditTemplate(user, command)
	handler := adaptor.HTTPHandler(templ.Handler(create))
	return handler(c)
}

func (u *MainHandler) HandleQuit(c *fiber.Ctx) error {
	c.ClearCookie("apiToken")
	return c.Redirect("/main")
}
