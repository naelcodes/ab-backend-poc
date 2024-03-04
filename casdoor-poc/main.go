package main

import (
	"fmt"
	"log/slog"

	"casdoor-poc/sdk"
	"casdoor-poc/views"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		slog.Error(fmt.Sprintf("error loading .env file - err : %v", err))
	}

	slog.Info("Loading environment loaded")
	sdk.InitCasdoorClient()
}

func main() {

	app := fiber.New()

	views.SetViews(app)

	app.Post("api/v1/redirect", func(c *fiber.Ctx) error {
		fmt.Println(fmt.Printf("POST code : %v , state : %v", c.Query("code"), c.Query("state")))
		return nil
	})

	app.Get("api/v1/redirect", func(c *fiber.Ctx) error {
		fmt.Println(fmt.Printf("GET code : %v , state : %v", c.Query("code"), c.Query("state")))
		return nil
	})

	app.Post("api/v1/signUp", func(c *fiber.Ctx) error {
		return nil
	})

	users, _ := sdk.CasdoorClient.GetUsers()
	slog.Info(fmt.Sprintf("users : %v", len(users)))
	app.Listen("localhost:3000")

}
