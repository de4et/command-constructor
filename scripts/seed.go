package main

import (
	"context"
	"fmt"
	"log"

	"github.com/de4et/command-constructor/db"
	"github.com/de4et/command-constructor/db/fixtures"
	"github.com/de4et/command-constructor/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "command-constructor"

func main() {
	// init db
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(dbname).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	// setup store
	store := &db.Store{
		User:    db.NewMongoUserStore(client, dbname),
		Command: db.NewMongoCommandStore(client, dbname),
	}

	user := fixtures.AddUser(store, "michael")
	fmt.Println(user)

	var params = types.CreateCommandTemplateParams{
		Name:        "send files via ssh",
		Description: "send files by pscp(Putty)",
		CommandName: "pscp",
		TemplateParams: []types.CommandParam{
			{
				// pscp -i "%USERPROFILE%/Documents/prin.ppk" -r ./bin root@77.232.42.104:/root/
				Name:         "",
				Description:  "Path: example -- root@127.0.0.1:/root/",
				Type:         types.TypeNameless,
				Value:        []string{},
				DefaultValue: "",
			},
			{
				Name:         "-i",
				Description:  "private key to send without authentication\nexample -- %USERPROFILE%/Documents/prin.ppk",
				Type:         types.TypeString,
				Value:        []string{},
				DefaultValue: "",
			},
			{
				Name:         "-r",
				Description:  "for sending directory\nexample -- ./bin",
				Type:         types.TypeString,
				Value:        []string{},
				DefaultValue: "",
			},
		},
	}
	command := fixtures.AddCommand(store, user, params)
	fmt.Println("commands:", command)
}
