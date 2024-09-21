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

	mongoStore := db.NewMongoUserStore(client, dbname)
	userHandler := api.NewUserHandler(mongoStore)

	apiv1 := app.Group("/api/v1")
	apiv1.Post("/user", userHandler.HandleCreateUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	app.Listen(":" + *listenAddr)
}
