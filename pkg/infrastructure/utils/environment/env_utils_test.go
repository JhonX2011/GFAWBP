package environment

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type envScenery struct {
	aResult  any
	key      string
	fallback string
	aError   error
}

type MockEnvLookup struct{}

func NewMockEnvLookup() IEnvLookup {
	return MockEnvLookup{}
}

func (m MockEnvLookup) LookupEnv(_ string) (string, bool) {
	return "", false
}

type MockOsSetEnv struct{}

func NewMockNewSetEnv() ISetEnv {
	return MockOsSetEnv{}
}

func (m MockOsSetEnv) Setenv(_ string, _ string) error {
	return assert.AnError
}

func givenEnvScenery(_ *testing.T) *envScenery {
	return &envScenery{}
}

func (g *envScenery) givenFallback(fallback string) {
	g.key = "test"
	g.fallback = fallback
}

func (g *envScenery) givenFallbackUnique(fallback string) {
	g.key = "test-xxx"
	g.fallback = fallback
}

func (g *envScenery) givenEnvironment() {
	LoadEnvironment()
}

func (g *envScenery) whenCheckOrAssignEnvIsExecuted() {
	checkOrAssignEnv(g.key, g.fallback, NewEnvLookup(), NewSetEnv())
}

func (g *envScenery) whenAssignEnvIsExecuted(v int) {
	assignEnv(g.key, g.fallback, v, NewSetEnv())
}

func (g *envScenery) whenIsFlagEnabledExecuted() {
	g.aResult = IsFlagEnabled(g.key)
}

func (g *envScenery) whenGetIntFeatureExecuted() {
	g.aResult = GetIntFeature(g.key)
}

func (g *envScenery) whenGetStringFeatureExecuted() {
	g.aResult = GetStringFeature(g.key)
}

func (g *envScenery) whenGetEncodeStringFeatureExecuted() {
	g.aResult, g.aError = GetEncodeStringFeature(g.key)
}

func (g *envScenery) whenPrintEnvExecuted() {
	PrintEnv()
}

func (g *envScenery) thenNoError(t *testing.T) {
	assert.Nil(t, g.aError)
}

func (g *envScenery) thenEqual(t *testing.T, expectedMessage any, output any) {
	assert.Equal(t, expectedMessage, output)
}

func (g *envScenery) thenHaveError(t *testing.T) {
	assert.NotNil(t, g.aError)
}

func (g *envScenery) thencheckOrAssignEnvWithLookupEnvFail(t *testing.T, envLookup IEnvLookup, setEnv ISetEnv) {
	assert.Panics(t, func() {
		checkOrAssignEnv(g.key, g.fallback, envLookup, setEnv)
	}, "A panic was expected due to an error in executing os.LookupEnv")
}

func (g *envScenery) thenAssignEnvFailExecuted(t *testing.T, v int) {
	assert.Panics(t, func() {
		assignEnv(g.key, "", v, NewSetEnv())
	}, "A panic was expected due to an error in executing assignEnv with value fallback zero")
}

func (g *envScenery) thenAssignEnvOk(t *testing.T, v int, setEnv ISetEnv) {
	assert.Panics(t, func() {
		assignEnv(g.key, g.fallback, v, setEnv)
	}, "panic was expected")
}

func (g *envScenery) thenGetIntFeatureFailExecuted(t *testing.T) {
	assert.Panics(t, func() {
		GetIntFeature(g.key)
	}, "A panic was expected due to a non-numeric value")
}

func TestCheckOrAssignEnvOk(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenEnvironment()
	e.key = "key-no-exist"
	e.fallback = "fallback key no exist"
	e.whenCheckOrAssignEnvIsExecuted()
	e.thenNoError(t)
}

//nolint:gocritic
func TestCheckOrAssignEnvWithLookupEnvFail(t *testing.T) {
	e := givenEnvScenery(t)
	/* monkey.Patch(os.LookupEnv, func(key string) (string, bool) {
		return "", false
	})
	defer monkey.Unpatch(os.LookupEnv)
	monkey.Patch(os.Setenv, func(key, value string) error {
		return assert.AnError
	})
	defer monkey.Unpatch(os.Setenv)*/

	e.key = "key-no-exist"
	e.fallback = "fallback key no exist"
	e.thencheckOrAssignEnvWithLookupEnvFail(t, NewMockEnvLookup(), NewMockNewSetEnv())
}

func TestCheckOrAssignEnvWithLookupEnvFailAndGetenvEmpty(t *testing.T) {
	e := givenEnvScenery(t)
	e.key = "key-test-test"
	e.fallback = ""
	e.whenAssignEnvIsExecuted(0)
	e.thencheckOrAssignEnvWithLookupEnvFail(t, NewMockEnvLookup(), NewSetEnv())
}

func TestAssignEnvOk(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallback("testValor")
	e.whenAssignEnvIsExecuted(0)
	e.thenNoError(t)
	e.thenEqual(t, "testValor", os.Getenv("test"))
}

func TestAssignEnvOkWithValidate(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallback("testValor")
	e.whenAssignEnvIsExecuted(1)
	e.thenNoError(t)
	e.thenEqual(t, "testValor", os.Getenv("test"))
}

func TestAssignEnvFailExecuted(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallbackUnique("testValor")
	e.whenAssignEnvIsExecuted(1)
	e.thenAssignEnvFailExecuted(t, 1)
}

func TestAssignEnvFailExecutedSetenvFailure(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallbackUnique("testValor")
	e.thenAssignEnvOk(t, 1, NewMockNewSetEnv())
}

func TestIsFlagEnabledTrue(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallback("testValor")
	e.whenAssignEnvIsExecuted(0)
	e.whenIsFlagEnabledExecuted()
	e.thenNoError(t)
	e.thenEqual(t, e.aResult, true)
}

func TestIsFlagEnabledFalse(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallback("")
	e.whenAssignEnvIsExecuted(0)
	e.whenIsFlagEnabledExecuted()
	e.thenNoError(t)
	e.thenEqual(t, e.aResult, false)
}

func TestGetIntFeatureOk(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallback("123")
	e.whenAssignEnvIsExecuted(0)
	e.whenGetIntFeatureExecuted()
	e.thenNoError(t)
	e.thenEqual(t, e.aResult, 123)
}

func TestGetIntFeatureFail(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallback("abc")
	e.whenAssignEnvIsExecuted(0)
	e.thenGetIntFeatureFailExecuted(t)
}

func TestGetStringFeatureOk(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallback("stringTest")
	e.whenAssignEnvIsExecuted(0)
	e.whenGetStringFeatureExecuted()
	e.thenNoError(t)
	e.thenEqual(t, e.aResult, os.Getenv("test"))
}

func TestGetStringDefaultOk(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallback("")
	e.whenAssignEnvIsExecuted(0)
	e.whenGetStringFeatureExecuted()
	e.thenNoError(t)
	e.thenEqual(t, e.aResult, os.Getenv("test"))
}

func TestGetEncodeStringFeatureOk(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallback(base64.StdEncoding.EncodeToString([]byte("encriptado")))
	e.whenAssignEnvIsExecuted(0)
	e.whenGetEncodeStringFeatureExecuted()
	e.thenNoError(t)
}

func TestGetEncodeStringFeatureFail(t *testing.T) {
	e := givenEnvScenery(t)
	e.givenFallback("eyJyZXNwb25zZSI6InRlc3RfdmFsIn0")
	e.whenAssignEnvIsExecuted(0)
	e.whenGetEncodeStringFeatureExecuted()
	e.thenHaveError(t)
}

func TestLoadEnvironment(t *testing.T) {
	e := givenEnvScenery(t)
	t.Setenv("ENVIRONMENT", "prod")
	t.Setenv("APP_NAME", "fury_fbm-invcontrol-dispatcher")
	t.Setenv("PORT", "8081")
	t.Setenv("STACK", "Go")
	t.Setenv("MODE_DEBUG", "0")
	t.Setenv("MAX_SEG_CACHE", "1500")
	t.Setenv("CACHE_ENABLED", "0")

	e.givenEnvironment()
	e.thenEqual(t, "prod", os.Getenv("ENVIRONMENT"))
	e.thenEqual(t, "fury_fbm-invcontrol-dispatcher", os.Getenv("APP_NAME"))
	e.thenEqual(t, "8081", os.Getenv("PORT"))
	e.thenEqual(t, "Go", os.Getenv("STACK"))
	e.thenEqual(t, "0", os.Getenv("MODE_DEBUG"))
	e.thenEqual(t, "1500", os.Getenv("MAX_SEG_CACHE"))
	e.thenEqual(t, "0", os.Getenv("CACHE_ENABLED"))
}

func TestLoadEnvironmentLocal(t *testing.T) {
	e := givenEnvScenery(t)
	t.Setenv("ENVIRONMENT", "local")
	t.Setenv("APP_NAME", "fury_fbm-invcontrol-dispatcher")
	t.Setenv("PORT", "8081")
	t.Setenv("STACK", "Go")
	t.Setenv("MODE_DEBUG", "0")
	t.Setenv("MAX_SEG_CACHE", "1500")
	t.Setenv("CACHE_ENABLED", "0")

	e.givenEnvironment()
	e.thenEqual(t, "local", os.Getenv("ENVIRONMENT"))
	e.thenEqual(t, "fury_fbm-invcontrol-dispatcher", os.Getenv("APP_NAME"))
	e.thenEqual(t, "8081", os.Getenv("PORT"))
	e.thenEqual(t, "Go", os.Getenv("STACK"))
	e.thenEqual(t, "0", os.Getenv("MODE_DEBUG"))
	e.thenEqual(t, "1500", os.Getenv("MAX_SEG_CACHE"))
	e.thenEqual(t, "0", os.Getenv("CACHE_ENABLED"))
}

func TestPrintEnv(t *testing.T) {
	e := givenEnvScenery(t)
	e.whenPrintEnvExecuted()
	e.thenNoError(t)
}
