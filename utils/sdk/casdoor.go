package sdk

import (
	"fmt"
	"os"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"neema.co.za/rest/utils/logger"
)

type CasdoorSDK struct {
	*casdoorsdk.Client
}

func NewCasdoorSdk() *CasdoorSDK {
	return &CasdoorSDK{GetCasdoorSdk()}
}

func GetCasdoorSdk() *casdoorsdk.Client {
	certificate, err := os.ReadFile("./certs/casdoor-certificate.pem")
	if err != nil {
		panic(fmt.Errorf("couldn't read the perm file : %v", err))
	}
	logger.Info("SERVER_URL : " + os.Getenv("CASDOOR_SERVER_URL"))
	logger.Info("CLIENT_ID : " + os.Getenv("CASDOOR_CLIENT_ID"))
	logger.Info("CLIENT_SECRET : " + os.Getenv("CASDOOR_CLIENT_SECRET"))
	logger.Info("ORGANIZATION_NAME : " + os.Getenv("CASDOOR_ORGANIZATION_NAME"))
	logger.Info("APPLICATION_NAME : " + os.Getenv("CASDOOR_APPLICATION_NAME"))

	return casdoorsdk.NewClient(os.Getenv("CASDOOR_SERVER_URL"), os.Getenv("CASDOOR_CLIENT_ID"), os.Getenv("CASDOOR_CLIENT_SECRET"), string(certificate), os.Getenv("CASDOOR_ORGANIZATION_NAME"), os.Getenv("CASDOOR_APPLICATION_NAME"))
}
