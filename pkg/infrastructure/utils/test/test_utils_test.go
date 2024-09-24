package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testUtilsScenery struct {
	aResult interface{}
	aError  error
	server  *httptest.Server
}

func givenUtilsScenery() *testUtilsScenery {
	return &testUtilsScenery{}
}

func (tu *testUtilsScenery) givenConfigRouteServer(handler http.HandlerFunc) {
	tu.server = httptest.NewServer(handler)
}

func (tu *testUtilsScenery) whenExecuteRequest(url string, body io.Reader, methodHTTP string) {
	url2 := fmt.Sprintf("%s/%s", tu.server.URL, url)
	aResult, aError := ExecuteRequest(url2, body, methodHTTP)
	tu.aResult, tu.aError = aResult, aError
	defer aResult.Body.Close() //nolint:gocritic
}

func (tu *testUtilsScenery) thenNoError(t *testing.T, statusCode int) {
	assert.Nil(t, tu.aError)
	assert.NotNil(t, tu.aResult)
	assert.Equal(t, statusCode, tu.aResult.(*http.Response).StatusCode)
}

func (tu *testUtilsScenery) thenError(t *testing.T) {
	assert.Nil(t, tu.aError)
	assert.NotNil(t, tu.aResult)
	assert.Equal(t, http.StatusInternalServerError, tu.aResult.(*http.Response).StatusCode)
}

func (tu *testUtilsScenery) thenClose() {
	tu.server.Close()
	if tu.aResult.(*http.Response) != nil {
		tu.aResult.(*http.Response).Body.Close()
	}
}

func TestRefreshConfigResponseFail(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusInternalServerError) })
	s := givenUtilsScenery()
	s.givenConfigRouteServer(testHandler)
	s.whenExecuteRequest("test", nil, "POST")
	s.thenError(t)
	s.thenClose()
}

func TestRefreshConfigResponseOk(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	s := givenUtilsScenery()
	s.givenConfigRouteServer(testHandler)
	s.whenExecuteRequest("/fbm/invcontrol/v1/", nil, "GET")
	s.thenNoError(t, http.StatusOK)
	s.thenClose()
}
