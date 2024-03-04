package sdk

import (
	"fmt"
	"log/slog"
	"os"

	CasdoorSDK "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

var CasdoorClient *CasdoorSDK.Client

func InitCasdoorClient() {

	certificate, err := os.ReadFile("./sdk/certificate.pem")

	if err != nil {
		panic(fmt.Errorf("couldn't read the perm file : %v", err))
	}
	slog.Info("SERVER_URL : " + os.Getenv("SERVER_URL"))
	slog.Info("CLIENT_ID : " + os.Getenv("CLIENT_ID"))
	slog.Info("CLIENT_SECRET : " + os.Getenv("CLIENT_SECRET"))
	slog.Info("ORGANIZATION_NAME : " + os.Getenv("ORGANIZATION_NAME"))
	slog.Info("APPLICATION_NAME : " + os.Getenv("APPLICATION_NAME"))

	CasdoorClient = CasdoorSDK.NewClient(os.Getenv("SERVER_URL"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), string(certificate), os.Getenv("ORGANIZATION_NAME"), os.Getenv("APPLICATION_NAME"))
}
