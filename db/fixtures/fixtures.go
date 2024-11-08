package fixtures

import (
	"context"
	"fmt"
	"log"

	"github.com/de4et/command-constructor/db"
	"github.com/de4et/command-constructor/types"
)

func AddUser(store *db.Store, name string) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Name:     name,
		Email:    fmt.Sprintf("%s@gmail.com", name),
		Password: fmt.Sprintf("%s_123!", name),
	})
	if err != nil {
		log.Fatal(err)
	}
	insertedUser, err := store.User.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}

func AddCommand(store *db.Store, user *types.User, params types.CreateCommandTemplateParams) *types.CommandTemplate {
	command, err := types.NewCommandTemplateFromParams(user.ID, params)
	if err != nil {
		log.Fatal(err)
	}

	insertedCommand, err := store.Command.InsertCommand(context.TODO(), command)
	if err != nil {
		log.Fatal(err)
	}
	return insertedCommand
}
