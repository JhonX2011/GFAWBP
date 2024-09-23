package utils

import "fmt"

const (
	Host              = "localhost"
	Protocol          = "http"
	AppPort           = "8080"
	MockServerAddress = ":9999"
)

func GetBaseURL() string {
	return fmt.Sprintf("%s://%s:%s", Protocol, Host, AppPort)
}
