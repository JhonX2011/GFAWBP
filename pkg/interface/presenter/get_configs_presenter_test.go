package presenter

import (
	"testing"

	mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"
	gt "github.com/JhonX2011/GFAWBP/test/generic"
)

type getConfigsPresenterScenery struct {
	getConfigsPresenter IGetConfigsPresenter
	gt.GenericTest
}

func givenGetConfigsPresenterScenery(t *testing.T) *getConfigsPresenterScenery {
	t.Parallel()
	s := getConfigsPresenterScenery{
		getConfigsPresenter: NewGetConfigsPresenter(),
	}

	return &s
}

func (f *getConfigsPresenterScenery) whenResponseGetConfigsIsCall() {
	f.AResult = f.getConfigsPresenter.ResponseGetConfigs([]mcs.ConfigMember{})
}

func TestResponseGetConfigsOK(t *testing.T) {
	s := givenGetConfigsPresenterScenery(t)
	s.whenResponseGetConfigsIsCall()
	s.ThenNotEmpty(t)
}
