package cross

import (
	"github.com/JhonX2011/GFAWBP/test/functional/rest_client"
	"github.com/JhonX2011/GOFunctionalTestsMocker/pkg/mock"
)

type FeatureCrossFunctions struct {
	RequestID  string
	Mocker     mock.Mocker
	RestClient restclient.IClient
}
