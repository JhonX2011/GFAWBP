package main

import (
	"context"
	"testing"

	restclient "github.com/JhonX2011/GFAWBP/pkg/test/functional/rest_client"
	"github.com/JhonX2011/GFAWBP/pkg/test/functional/steps/default"
	"github.com/JhonX2011/GFAWBP/pkg/test/functional/steps/ping"
	"github.com/JhonX2011/GFAWBP/pkg/test/functional/utils"
	"github.com/JhonX2011/GOFunctionalTestsMocker/pkg/mock"
	"github.com/cucumber/godog"
)

const (
	pingSceneryPath = "../../test/functional/features/ping"
	mainSceneryPath = "../../test/functional/features/default"
)

func TestSuites(t *testing.T) {
	//Load environment variables if necessary
	go func() {
		err := run()
		if err != nil {
			panic(err)
		}
	}()
	err := utils.WaitServerRunning()
	if err != nil {
		t.Errorf("failed awaiting server up when runinng test suites")
		t.Fail()
		return
	}
	mockServer, vmock := mock.New()
	runMockServer(t, err, mockServer)
	suites := buildTestSuites(vmock)
	for _, suite := range suites {
		if suite.Run() != 0 {
			t.Fail()
		}
	}
}

func runMockServer(t *testing.T, err error, mockServer mock.Router) {
	go func() {
		err = mockServer.Run(utils.MockServerAddress)
		if err != nil {
			t.Error(err)
			t.Fail()
			panic(err)
		}
	}()
}

func buildTestSuites(mocker mock.Mocker) []godog.TestSuite {
	var suites []godog.TestSuite
	suites = append(suites,
		buildSuite("Ping tests", pingSceneryInitializer, pingSceneryPath),
		//buildSuite("Default tests", defaultSceneryInitializer(mocker), mainSceneryPath),
	)
	return suites
}

func buildSuite(suiteName string, initializer func(*godog.ScenarioContext), feature string) godog.TestSuite {
	return godog.TestSuite{
		Name:                suiteName,
		ScenarioInitializer: initializer,
		Options:             buildOptions(feature),
	}
}

func buildOptions(paths ...string) *godog.Options {
	return &godog.Options{
		Format:      "pretty",
		Paths:       paths,
		Tags:        "~@wip",
		Randomize:   0, // randomize scenario execution order
		Strict:      true,
		Concurrency: 0,
	}
}

func defaultSceneryInitializer(mocker mock.Mocker) func(*godog.ScenarioContext) {
	return func(s *godog.ScenarioContext) {
		//configService, _ := configs.NewConfigService()
		restClient := restclient.New()
		featureBaseFunctions := defaultSteps.NewDefaultFeatureFunctions(s, restClient, mocker)

		s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
			featureBaseFunctions.Reset()
			return ctx, nil
		})
	}
}

func pingSceneryInitializer(s *godog.ScenarioContext) {
	restClient := restclient.New()
	pingFeature := ping.NewPingFeature(s, restClient)

	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		pingFeature.Reset()
		return ctx, nil
	})
}
