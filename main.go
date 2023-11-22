// main.go
package main

import (
	"log"

	"neema.co.za/rest/ioc"
	userModule "neema.co.za/rest/modules/user"
	App "neema.co.za/rest/utils/app"
)

func main() {

	app := App.Initialise()

	defer func() { // Start the Fiber app
		log.Fatal(app.Listen(":8080"))
	}()

	userModule.Start(ioc.InjectUserModule())
}
