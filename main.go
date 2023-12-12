// main.go
package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	logger "neema.co.za/rest/utils/logger"

	customerModule "neema.co.za/rest/modules/customer"
	App "neema.co.za/rest/utils/app"
)

func init() {
	logger.Info("Loading environment variables")

	if err := godotenv.Load(); err != nil {
		logger.Error(fmt.Sprintf("error loading .env file - err : %v", err))
	}

	logger.Info("Loading environment loaded")

}

func main() {

	app := App.Initialise()

	defer app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))

	routerV1 := app.Group(os.Getenv("API_V1_BASE_PATH"))

	routerV1.Mount("/customers", customerModule.GetModule().App)

}
