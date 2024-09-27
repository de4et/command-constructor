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
	command := fixtures.AddCommand(store, user, "scp", "sending files via ssh",
		[]types.CommandParam{
			{
				Name:        "-r",
				Description: "for sending entire directories",
			},
		},
		[]types.CommandParam{
			{
				Name:        "-v",
				Description: "for Debug messages",
			},
		},
	)
	fmt.Println(command)
}
