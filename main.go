package main

import (
	"context"
	"flag"
	"log"

	"github.com/de4et/command-constructor/api"
	"github.com/de4et/command-constructor/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "command-constructor"

var config = fiber.Config{
	ErrorHandler: api.ErrorHandler,
	Views:        api.GetEngine(),
}

// config
// logger
// design + html
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("port", "5000", "Listen address of API")
	flag.Parse()

	app := fiber.New(config)

	store := &db.Store{
		User:    db.NewMongoUserStore(client, dbname),
		Command: db.NewMongoCommandStore(client, dbname),
	}
	api.SetupRoutes(app, store)

	app.Listen(":" + *listenAddr)
}

// TODO: finish README
