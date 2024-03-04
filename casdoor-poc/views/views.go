package views

import (
	components "casdoor-poc/views/components/root"
	pages "casdoor-poc/views/pages"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func SetViews(app *fiber.App) {
	app.Get("/", Render(pages.HomePage()))
	app.Get("/login", Render(pages.LoginPage()))
	app.Get("/signup", Render(pages.SignUpPage()))
}

func Render(content templ.Component) fiber.Handler {
	return adaptor.HTTPHandler(templ.Handler(components.Root(content)))
}
