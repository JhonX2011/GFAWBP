package cross

import (
	restclient "github.com/JhonX2011/GFAWBP/pkg/test/functional/rest_client"
	"github.com/JhonX2011/GOFunctionalTestsMocker/pkg/mock"
)

type FeatureCrossFunctions struct {
	RequestID  string
	Mocker     mock.Mocker
	RestClient restclient.IClient
}
