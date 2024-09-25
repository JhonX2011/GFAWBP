package defaultsteps

import (
	"github.com/JhonX2011/GFAWBP/test/functional/rest_client"
	"github.com/JhonX2011/GFAWBP/test/functional/steps/cross_step"
	"github.com/JhonX2011/GFAWBP/test/functional/utils"
	"github.com/JhonX2011/GOFunctionalTestsMocker/pkg/mock"
	"github.com/cucumber/godog"
)

type DefaultFeatureFunctions struct {
	cross.FeatureCrossFunctions
}

func NewDefaultFeatureFunctions(s *godog.ScenarioContext, restClient restclient.IClient, mocker mock.Mocker) *DefaultFeatureFunctions {
	crossFeatureFunctions := cross.FeatureCrossFunctions{
		RequestID:  utils.GenerateUUID(),
		Mocker:     mocker,
		RestClient: restClient,
	}
	defaultFeatureFunctions := &DefaultFeatureFunctions{
		FeatureCrossFunctions: crossFeatureFunctions,
	}
	return defaultFeatureFunctions
}

func loadSteps(s *godog.ScenarioContext, f *DefaultFeatureFunctions) {
	s.Step(`^the cities from country "([^"]*)" are:$`, f.validateCitiesByCountry)
}

func (f *DefaultFeatureFunctions) validateCitiesByCountry(CountryID string, table *godog.Table) error {
	//origin := ""
	//inventory := ""
	//for _, row := range table.Rows[1:] {
	//	var responseBody models.RulesByOrigin
	//	origin = row.Cells[0].Value
	//	inventory = row.Cells[1].Value
	//	compatiblesFCStrings := strings.Split(row.Cells[2].Value, ",")
	//	typology := row.Cells[3].Value
	//
	//	costStrings := strings.Split(row.Cells[4].Value, ",")
	//
	//	for i, compatiblesFC := range compatiblesFCStrings {
	//		cost := parseCost(costStrings[i])
	//		edge := &models.Edge{
	//			Origin: models.Node{ID: origin},
	//			Destination: models.Node{
	//				ID: compatiblesFC,
	//			},
	//			TransportCharacteristics: models.TransportCharacteristics{
	//				Cost: cost,
	//			},
	//		}
	//
	//		responseBody.Edges = append(responseBody.Edges, edge)
	//	}
	//
	//	responseBody.Inventory = models.InventoryDetails{
	//		Typologies: strings.Split(typology, ","),
	//	}
	//
	//	params := map[string]string{
	//		"use-case":     useCaseRC,
	//		"origin-node":  origin,
	//		"inventory-id": inventory,
	//	}
	//
	//	f.Cases.RequestPre.NodeID = origin
	//	responseBodyFinal, err := json.Marshal(responseBody)
	//	if err != nil {
	//		return err
	//	}
	//
	//	header := map[string]string{
	//		cross.RequestID: f.RequestID,
	//	}
	//
	//	url := fmt.Sprintf(rulesByOriginURL, siteID)
	//	err = mockserver.GenerateMockResource(
	//		f.Mocker,
	//		http.MethodGet,
	//		http.StatusOK,
	//		"",
	//		responseBodyFinal,
	//		url,
	//		params,
	//		header,
	//	)
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (f *DefaultFeatureFunctions) Reset() {
	f.RequestID = utils.GenerateUUID()
}
