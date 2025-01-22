package api

import (
	"net/http"

	"github.com/de4et/command-constructor/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/gofiber/template/html/v2"
)

func SetupRoutes(app *fiber.App, store *db.Store) {
	userHandler := NewUserHandler(store)
	commandHandler := NewCommandHandler(store)
	mainHandler := NewMainHandler(store)

	// API
	// auth and reg
	app.Post("/api/auth", userHandler.HandleAuthenticate)
	app.Post("/api/reg", userHandler.HandleCreateUser)

	// Versioned api
	apiv1 := app.Group("/api/v1", JWTAuth(store))
	// user handlers
	apiv1.Delete("/user", userHandler.HandleDeleteUser)
	apiv1.Get("/user", userHandler.HandleGetUser)

	// command handlers
	apiv1.Get("/command", commandHandler.HandleGetCommands)
	apiv1.Post("/command", commandHandler.HandleCreateCommand)
	apiv1.Delete("/command/:id", commandHandler.HandleDeleteCommand)
	apiv1.Get("/command/search/:name", commandHandler.HandleSearchCommands)
	apiv1.Put("/command/:id", commandHandler.HandleUpdateCommand)

	// App
	app.Static("/static/", "static/")

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/main",
		},
		StatusCode: http.StatusMovedPermanently,
	}))

	app.Use(func(c *fiber.Ctx) error {
		if c.Protocol() == "http" {
			return c.Redirect("https://"+c.Hostname()+c.OriginalURL(), fiber.StatusMovedPermanently)
		}
		return c.Next()
	})

	app.Get("/main", AuthMiddleware(store), mainHandler.HandleMain)
	app.Get("/create", AuthMiddleware(store), mainHandler.HandleCreate)
	app.Get("/edit/:id", AuthMiddleware(store), mainHandler.HandleEdit)
	app.Get("/quit", AuthMiddleware(store), mainHandler.HandleQuit)
}

func GetEngine() (engine *html.Engine) {
	engine = html.New("./view", ".html")
	return
}
