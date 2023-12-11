package app

import (
	"github.com/gofiber/fiber/v2"
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
	app.Use(getRecover())
	return app
}
