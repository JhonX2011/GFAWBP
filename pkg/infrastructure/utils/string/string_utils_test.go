package stringu

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const UUIDKeyValue = "PreOptimization-398e2b3e-c6e0-11ee-8cd1-367ddaaad6e9"

type stringUtilsScenery struct {
	aResult any
	aError  error
}

func givenStringUtilsScenery(t *testing.T) *stringUtilsScenery {
	t.Parallel()
	return &stringUtilsScenery{}
}

func (s *stringUtilsScenery) whenGenerateUUIDIsExecuted() {
	s.aResult = GenerateUUID()
}

func (s *stringUtilsScenery) whenFindSubstringInSliceIsExecuted(sli []string, contain string) {
	s.aResult = FindSubstringInSlice(sli, contain)
}

func (s *stringUtilsScenery) whenContainsIsExecuted(arr []string, item string) {
	s.aResult = Contains(arr, item)
}

func (s *stringUtilsScenery) whenDecodeBase64ToStructIsExecuted(encodedBase64 string) {
	s.aResult, s.aError = DecodeBase64ToStruct[interface{}](encodedBase64)
}

func (s *stringUtilsScenery) whenFindInLogIsExecuted(logs []string, word string) {
	s.aResult = FindInLog(logs, word)
}

func (s *stringUtilsScenery) whenGetRequestKeyIsExecuted(ctx context.Context, key any) {
	s.aResult = GetRequestKey(ctx, key)
}

func (s *stringUtilsScenery) thenNotEmpty(t *testing.T) {
	assert.NotNil(t, s.aResult)
}

func (s *stringUtilsScenery) thenEqual(t *testing.T, result any) {
	assert.Equal(t, result, s.aResult)
}

func (s *stringUtilsScenery) thenHaveAError(t *testing.T) {
	assert.NotNil(t, s.aError)
}

func TestGenerateUUID(t *testing.T) {
	u := givenStringUtilsScenery(t)
	u.whenGenerateUUIDIsExecuted()
	u.thenNotEmpty(t)
}

func TestFindSubstringInSlice(t *testing.T) {
	u := givenStringUtilsScenery(t)
	u.whenFindSubstringInSliceIsExecuted([]string{"slice1", "slice2"}, "slice2")
	u.thenNotEmpty(t)
	u.thenEqual(t, "slice2")
}

func TestFindSubstringNotInSlice(t *testing.T) {
	u := givenStringUtilsScenery(t)
	u.whenFindSubstringInSliceIsExecuted([]string{"slice1", "slice2"}, "slice3")
	u.thenNotEmpty(t)
	u.thenEqual(t, "")
}

func TestContainsTrue(t *testing.T) {
	u := givenStringUtilsScenery(t)
	u.whenContainsIsExecuted([]string{"1", "2", "3"}, "2")
	u.thenEqual(t, true)
}

func TestContainsFalse(t *testing.T) {
	u := givenStringUtilsScenery(t)
	u.whenContainsIsExecuted([]string{"1", "2", "3"}, "8")
	u.thenEqual(t, false)
}

func TestFindInLog(t *testing.T) {
	u := givenStringUtilsScenery(t)
	u.whenFindInLogIsExecuted([]string{"log1", "log2"}, "log1")
	u.thenEqual(t, true)
}

func TestFindInLogFalse(t *testing.T) {
	u := givenStringUtilsScenery(t)
	u.whenFindInLogIsExecuted([]string{"log1", "log2"}, "log3")
	u.thenEqual(t, false)
}

func TestDecodeBase64ToStruct(t *testing.T) {
	u := givenStringUtilsScenery(t)
	u.whenDecodeBase64ToStructIsExecuted("e30=")
	u.thenEqual(t, map[string]interface{}{})
}

func TestDecodeBase64ToStructFail(t *testing.T) {
	u := givenStringUtilsScenery(t)
	u.whenDecodeBase64ToStructIsExecuted("invalid_base64")
	u.thenHaveAError(t)
}

func TestDecodeBase64ToStructFailUnMarshall(t *testing.T) {
	u := givenStringUtilsScenery(t)
	u.whenDecodeBase64ToStructIsExecuted("aW52YWxpZF9qc29u")
	u.thenHaveAError(t)
}

func TestGetRequestKeyOK(t *testing.T) {
	u := givenStringUtilsScenery(t)
	ctx := context.WithValue(context.Background(), RequestKey, UUIDKeyValue)
	u.whenGetRequestKeyIsExecuted(ctx, RequestKey)
	u.thenEqual(t, UUIDKeyValue)
}

func TestGetRequestKeyFail(t *testing.T) {
	u := givenStringUtilsScenery(t)
	ctx := context.WithValue(context.Background(), RequestKey, UUIDKeyValue)
	u.whenGetRequestKeyIsExecuted(ctx, "AnyRequestKey")
	u.thenEqual(t, "")
}
