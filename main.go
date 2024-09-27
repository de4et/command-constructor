package main

import (
	"context"
	"flag"
	"log"

	"github.com/de4et/command-constructor/api"
	"github.com/de4et/command-constructor/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "command-constructor"

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

// TODO:
// user
// command
// config
// logger
// db(mongo or postgres or sqlite) -> mongo
// design + html
func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("port", "5000", "Listen address of API")
	flag.Parse()

	app := fiber.New(config)

	mongoUserStore := db.NewMongoUserStore(client, dbname)
	mongoCommandStore := db.NewMongoCommandStore(client, dbname)
	store := &db.Store{
		User:    mongoUserStore,
		Command: mongoCommandStore,
	}
	userHandler := api.NewUserHandler(store)
	commandHandler := api.NewCommandHandler(store)

	apiv1 := app.Group("/api/v1", api.JWTAuth)

	// auth
	app.Post("/api/auth", userHandler.HandleAuthenticate)

	// user handlers
	apiv1.Post("/user", userHandler.HandleCreateUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	// command handlers
	apiv1.Get("/command", commandHandler.HandleGetCommands)
	apiv1.Post("/command", commandHandler.HandleCreateCommand)
	apiv1.Get("/command/search/:name", commandHandler.HandleSearchCommand)
	apiv1.Put("/command/:id", commandHandler.HandleUpdateCommand)
	app.Listen(":" + *listenAddr)
}
