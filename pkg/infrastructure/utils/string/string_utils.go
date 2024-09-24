package stringu

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

type contextKey string

const (
	RequestKey contextKey = "requestID"
)

func GenerateUUID() string {
	newUUID, _ := uuid.NewUUID()
	return newUUID.String()
}

func FindSubstringInSlice(slice []string, contain string) string {
	for _, elem := range slice {
		if strings.Contains(elem, contain) {
			return elem
		}
	}

	return ""
}

func Contains(s []string, item string) bool {
	m := ConvertSliceToMap(s)
	_, ok := m[item]
	return ok
}

func ConvertSliceToMap(arr []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, item := range arr {
		m[item] = struct{}{}
	}
	return m
}

func FindInLog(logs []string, w string) bool {
	wordsMap := make(map[string]bool)

	for _, log := range logs {
		words := strings.Fields(log)
		for _, word := range words {
			wordsMap[word] = true
		}
	}

	return wordsMap[w]
}

func DecodeBase64ToStruct[T any](encodedBase64 string) (T, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedBase64)
	if err != nil {
		return *new(T), err
	}

	var targetStruct T
	if errJSON := json.Unmarshal(decodedBytes, &targetStruct); errJSON != nil {
		return *new(T), errJSON
	}

	return targetStruct, nil
}

func GetRequestKey(ctx context.Context, key any) string {
	val, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return val
}
