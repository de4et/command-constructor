package api

import (
	"fmt"

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
	user := c.Context().Value("user").(*types.User)
	if user == nil {
		return fmt.Errorf("no user")
	}

	commands, err := u.Store.Command.GetCommandsOfUser(c.Context(), user.ID)
	if err != nil {
		return err
	}
	return c.JSON(commands)
}

func (u *CommandHandler) HandleDeleteCommand(c *fiber.Ctx) error {
	id := c.Params("id")

	user := c.Context().Value("user").(*types.User)
	if user == nil {
		return fmt.Errorf("no user")
	}

	if ok, err := u.Store.Command.IsExists(c.Context(), user.ID, id); err != nil {
		return ErrPrivateInternal()
	} else if !ok {
		return ErrNotExist()
	}

	if err := u.Store.Command.DeleteCommandByID(c.Context(), id); err != nil {
		return ErrPrivateInternal()
	}

	return c.JSON(map[string]string{"deleted": id})
}

func (u *CommandHandler) HandleCreateCommand(c *fiber.Ctx) error {
	var params types.CreateCommandTemplateParams
	err := c.BodyParser(&params)
	if err != nil {
		return ErrInvalidData()
	}

	user := c.Context().Value("user").(*types.User)
	if user == nil {
		return ErrUnauthorized()
	}

	cmd, err := types.NewCommandTemplateFromParams(user.ID, params)
	if err != nil {
		return ErrPrivateInternal()
	}

	insertedCmd, err := u.Store.Command.InsertCommand(c.Context(), cmd)
	if err != nil {
		return ErrPrivateInternal()
	}

	return c.JSON(insertedCmd)
}

func (u *CommandHandler) HandleSearchCommands(c *fiber.Ctx) error {
	name := c.Params("name")

	user := c.Context().Value("user").(*types.User)
	if user == nil {
		return fmt.Errorf("no user")
	}
	res, err := u.Store.Command.SearchCommands(c.Context(), user.ID, name)
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

	user := c.Context().Value("user").(*types.User)
	if user == nil {
		return fmt.Errorf("no user")
	}

	if ok, err := u.Store.Command.IsExists(c.Context(), user.ID, id); err != nil {
		return ErrPrivateInternal()
	} else if !ok {
		return ErrNotExist()
	}

	cmd, err := types.NewCommandTemplateFromParams(user.ID, params.CreateCommandTemplateParams)
	if err != nil {
		return err
	}

	err = u.Store.Command.UpdateCommand(c.Context(), id, cmd)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": user.ID})
}
