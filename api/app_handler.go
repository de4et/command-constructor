package api

import (
	"fmt"

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
	// c.Response().Header.Add("Cache-Control", "private, no-cache, no-store, must-revalidate")
	// c.Request().Header.Add("Cache-Control", "private, no-cache, no-store, must-revalidate")
	fmt.Println(c.Context().Value("user"))
	user := c.Context().Value("user").(*types.User)

	commandTemplates := make([]types.CommandTemplate, 0)
	// user := &types.User{
	// 	Name: "de4et",
	// }
	// var user *types.User
	index := view.Index(commandTemplates, user)
	handler := adaptor.HTTPHandler(templ.Handler(index))
	return handler(c)

	// API + Fetch(in js)
}

func (u *MainHandler) HandleQuit(c *fiber.Ctx) error {
	// TODO
	c.ClearCookie("apiToken")
	return c.Redirect("/main", 200)
}
