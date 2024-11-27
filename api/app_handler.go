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

	commandTemplates := make([]types.CommandTemplate, 0)

	index := view.Index(commandTemplates, user)
	handler := adaptor.HTTPHandler(templ.Handler(index))
	return handler(c)
}

func (u *MainHandler) HandleQuit(c *fiber.Ctx) error {
	// TODO
	c.ClearCookie("apiToken")
	return c.Redirect("/main")
}
