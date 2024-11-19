package api

import (
	"fmt"

	"github.com/de4et/command-constructor/db"
	"github.com/gofiber/fiber/v2"
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
	// return nil
	// c.Response().Header.Add("Cache-Control", "private, no-cache, no-store, must-revalidate")
	// c.Request().Header.Add("Cache-Control", "private, no-cache, no-store, must-revalidate")
	fmt.Println("im hereerererererer")
	return c.Render("index", nil)

	// API + Fetch(in js)
}
