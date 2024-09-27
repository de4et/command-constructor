package api

import (
	"github.com/de4et/command-constructor/db"
	"github.com/de4et/command-constructor/types"
	"github.com/gofiber/fiber/v2"
)

type CommandHandler struct {
	Store *db.Store
}

func NewCommandHandler(store *db.Store) *CommandHandler {
	return &CommandHandler{
		Store: store,
	}
}

func (u *CommandHandler) HandleGetCommands(c *fiber.Ctx) error {
	commands, err := u.Store.Command.GetCommands(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(commands)
}

func (u *CommandHandler) HandleCreateCommand(c *fiber.Ctx) error {
	var params types.CreateCommandTemplateParams
	err := c.BodyParser(&params)
	if err != nil {
		return err
	}

	const userID = "66f1612f7d88cdd84ddca7de" // FIXME
	cmd, err := types.NewCommandTemplateFromParams(userID, params)
	if err != nil {
		return err
	}

	insertedCmd, err := u.Store.Command.InsertCommand(c.Context(), cmd)
	if err != nil {
		return err
	}

	return c.JSON(insertedCmd)
}

func (u *CommandHandler) HandleSearchCommand(c *fiber.Ctx) error {
	name := c.Params("name")
	const userID = "66f1612f7d88cdd84ddca7de" // FIXME
	res, err := u.Store.Command.SearchCommand(c.Context(), userID, name)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (u *CommandHandler) HandleUpdateCommand(c *fiber.Ctx) error {
	id := c.Params("id")
	var params types.UpdateCommandTemplateParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	const userID = "66f1612f7d88cdd84ddca7de" // FIXME
	cmd, err := types.NewCommandTemplateFromParams(userID, params.CreateCommandTemplateParams)
	if err != nil {
		return err
	}

	err = u.Store.Command.UpdateCommand(c.Context(), id, cmd)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": userID})
}
