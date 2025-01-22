package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/de4et/command-constructor/api"
	"github.com/de4et/command-constructor/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://mongodb:27017"
const dbname = "command-constructor"

var config = fiber.Config{
	ErrorHandler: api.ErrorHandler,
	Views:        api.GetEngine(),
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	listenAddr := flag.String("port", "5000", "Listen address of API")
	flag.Parse()

	app := fiber.New(config)

	store := &db.Store{
		User:    db.NewMongoUserStore(client, dbname),
		Command: db.NewMongoCommandStore(client, dbname),
	}
	api.SetupRoutes(app, store)

	go func() {
		log.Fatal(app.Listen(":" + *listenAddr))
	}()

	log.Fatal(app.ListenTLS(":443", os.Getenv("SSL_CERT_PATH"), os.Getenv("SSL_KEY_PATH")))
}

// TODO: finish README
