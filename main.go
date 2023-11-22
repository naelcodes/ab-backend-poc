// main.go
package main

import (
	"log"

	App "neema.co.za/rest/app"
)

func main() {

	app := App.Initialise()

	injectUserModule().Start()

	// Start the Fiber app
	log.Fatal(app.Listen(":8080"))
}
