// main.go
package main

import (
	"log"

	"github.com/joho/godotenv"

	"neema.co.za/rest/ioc"
	userModule "neema.co.za/rest/modules/user"
	App "neema.co.za/rest/utils/app"
)

func init() {
	godotenv.Load()
}

func main() {

	app := App.Initialise()

	defer log.Fatal(app.Listen(":8080"))

	userModule.Start(ioc.InjectUserModule())
}
