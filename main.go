// main.go
package main

import (
	"log"

	"github.com/joho/godotenv"

	userModule "neema.co.za/rest/modules/user"
	App "neema.co.za/rest/utils/app"
)

func init() {
	godotenv.Load()
}

const API_V1_BASE_PATH = "/api/v1"

func main() {

	app := App.Initialise()
	defer log.Fatal(app.Listen(":8080"))

	routerV1 := app.Group(API_V1_BASE_PATH)

	routerV1.Mount("/customers", userModule.GetApp())
}
