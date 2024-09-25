package ping

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/JhonX2011/GFAWBP/test/functional/rest_client"
	"github.com/JhonX2011/GFAWBP/test/functional/utils"
	"github.com/cucumber/godog"
)

const path = "/ping"

type (
	FeaturePing struct {
		restClient restclient.IClient
		cases      *pingStep
	}

	pingStep struct {
		responseStatus int
		responseBody   []byte
	}
)

func NewPingFeature(s *godog.ScenarioContext, restClient restclient.IClient) *FeaturePing {
	f := &FeaturePing{
		restClient: restClient,
		cases:      &pingStep{},
	}
	loadSteps(s, f)
	return f
}

func (f *FeaturePing) getServerStatus() error {
	url := fmt.Sprintf("%s%s", utils.GetBaseURL(), path)
	response, err := f.restClient.DoGet(context.Background(), url) //nolint:bodyclose
	if err != nil {
		return err
	}
	f.cases.responseStatus = response.StatusCode
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	f.cases.responseBody = bodyBytes
	return nil
}

func (f *FeaturePing) theServerStatusCodeIs(status int) error {
	if f.cases.responseStatus != status {
		return fmt.Errorf("the server is not up and running")
	}
	return nil
}

func (f *FeaturePing) theServerIsRunning() error {
	if f.cases.responseStatus != http.StatusOK {
		return fmt.Errorf("the server is not up and running")
	}
	return nil
}

func (f *FeaturePing) assertBodyResponse(s string) error {
	if strings.Trim(string(f.cases.responseBody), "\"") != s {
		return fmt.Errorf("the response body is not equal")
	}
	return nil
}

func loadSteps(s *godog.ScenarioContext, f *FeaturePing) {
	s.Step(`^i verify if server is up and running`, f.getServerStatus)
	s.Step(`^the server is up and running`, f.theServerIsRunning)
	s.Step(`^the server status code is (\d+)$`, f.theServerStatusCodeIs)
	s.Step(`^the response is "([^"]*)"$`, f.assertBodyResponse)
}

func (f *FeaturePing) Reset() {
	f.cases = &pingStep{}
}
