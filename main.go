package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	logger "neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/managers"

	CustomerModule "neema.co.za/rest/modules/customer"
	TravelerModule "neema.co.za/rest/modules/traveler"
	App "neema.co.za/rest/utils/app"
)

func init() {
	logger.Info("Loading environment variables")

	if err := godotenv.Load(".env"); err != nil {
		logger.Error(fmt.Sprintf("error loading .env file - err : %v", err))
	}

	logger.Info("Loading environment loaded")

}

func main() {

	app := App.Initialise()

	defer app.Listen(fmt.Sprintf("localhost:%s", os.Getenv("APP_PORT")))

	routerV1 := app.Group(os.Getenv("API_V1_BASE_PATH"))

	dependencyManager := managers.NewDependencyManager()

	customerModule := CustomerModule.GetModule(dependencyManager)
	travelerModule := TravelerModule.GetModule(dependencyManager)

	logger.Info(fmt.Sprintf("Dependencies Count: %v", len(dependencyManager.GetAll())))

	//logger.Info(fmt.Sprintf("API V1 Base Path: %v", len(app.GetRoutes())))

	routerV1.Mount("/customers", customerModule.App)
	routerV1.Mount("/travelers", travelerModule.App)

}
