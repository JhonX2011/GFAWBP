package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"

	restclient "github.com/JhonX2011/GFAWBP/pkg/test/functional/rest_client"
	"github.com/google/uuid"
)

const (
	helpDeskPath = "/ping"
	retries      = 40
	retryWait    = 500 * time.Millisecond
)

func WaitServerRunning() error {
	client := restclient.New()
	url := fmt.Sprintf("%s%s", GetBaseURL(), helpDeskPath)
	for i := 0; i < retries; i++ {
		response, err := client.DoGet(context.Background(), url) //nolint:all
		if err == nil && response.StatusCode == http.StatusOK {
			return nil
		}
		time.Sleep(retryWait)
	}
	return fmt.Errorf("the server did not respond the ping after [%d] retries", retries)
}

func GenerateUUID() string {
	generatedUUID, _ := uuid.NewUUID()
	return generatedUUID.String()
}
