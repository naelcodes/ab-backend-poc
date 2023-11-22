package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type RouterCreator func(prefix string, handlers ...fiber.Handler) fiber.Router
type Router fiber.Router
type App struct {
	*fiber.App
}

var app *App

func init() {
	app = &App{new(fiber.App)}
}

func NewRouter() RouterCreator {
	return app.Group
}

func Initialise() (app *App) {
	app.Use(getCors())
	return app
}
func getCors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
	})
}
