package doubles

import (
	mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"
)

func GetLoadStructError(s, m string, code int, originalError error, isRetryable bool) *mcs.MessageErrorResponse {
	var objE mcs.MessageErrorResponse
	objE.Status = s
	objE.Message = m
	objE.Code = code
	objE.OriginalError = originalError
	objE.IsRetryable = isRetryable

	objE.ErrorsInfo = make([]mcs.ObjectErrors, 0)

	return &objE
}
