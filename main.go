// main.go
package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	logger "neema.co.za/rest/utils/logger"

	userModule "neema.co.za/rest/modules/user"
	App "neema.co.za/rest/utils/app"
)

func init() {
	logger.Info("Loading environment variables")

	if err := godotenv.Load(); err != nil {
		fmt.Println("error loading .env file - err :", err)
	}

	logger.Info("Loading environment loaded")

}

func main() {

	app := App.Initialise()
	defer app.Listen(":8080")

	routerV1 := app.Group(os.Getenv("API_V1_BASE_PATH"))

	routerV1.Mount("/users", userModule.GetModule().App)
	fmt.Println("hello")
}
