package presenter

import (
	"errors"
	"testing"

	mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"
	gt "github.com/JhonX2011/GFAWBP/pkg/test/generic"
)

type getStructErrorPresenterScenery struct {
	getStructErrorPresenter IGetStructErrorPresenter
	gt.GenericTest
}

func givenStructErrorPresenterScenery(t *testing.T) *getStructErrorPresenterScenery {
	t.Parallel()
	s := getStructErrorPresenterScenery{
		getStructErrorPresenter: NewGetGetStructErrorPresenter(),
	}

	return &s
}

func (f *getStructErrorPresenterScenery) whenLoadStructErrorIsCall(errorsInfo []mcs.ObjectErrors) {
	err := errors.New("some error")
	f.AResult = f.getStructErrorPresenter.LoadStructError("status", "message", 0, err, errorsInfo, IsRetryableFalse)
}

func TestLoadStructErrorOK(t *testing.T) {
	s := givenStructErrorPresenterScenery(t)
	s.whenLoadStructErrorIsCall([]mcs.ObjectErrors{})
	s.ThenNotEmpty(t)
}

func TestLoadStructErrorFail(t *testing.T) {
	s := givenStructErrorPresenterScenery(t)
	s.whenLoadStructErrorIsCall(nil)
	s.ThenNotEmpty(t)
}
