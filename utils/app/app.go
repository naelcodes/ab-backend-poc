package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// type RouterCreator func(prefix string, handlers ...fiber.Handler) fiber.Router
// type Router fiber.Router
type App struct {
	*fiber.App
}

var app *App

func init() {
	app = &App{NewFiberApp()}
}

func NewFiberApp() *fiber.App {
	return fiber.New()
}

//func NewRouter() RouterCreator {
//	return app.Group
//}

func Initialise() *App {
	app.Use(getCors())
	return app
}

func getCors() func(*fiber.Ctx) error {
	config := cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}

	return cors.New(config)
}
