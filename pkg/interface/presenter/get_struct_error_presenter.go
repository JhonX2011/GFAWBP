package presenter

import mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"

const (
	IsRetryableFalse = false
	IsRetryableTrue  = true
)

type getStructErrorPresenter struct{}

type IGetStructErrorPresenter interface {
	LoadStructError(string, string, int, error, []mcs.ObjectErrors, bool) *mcs.MessageErrorResponse
}

func NewGetGetStructErrorPresenter() IGetStructErrorPresenter {
	return &getStructErrorPresenter{}
}

func (g *getStructErrorPresenter) LoadStructError(s, m string, code int, originalError error,
	errorsInfo []mcs.ObjectErrors, isRetryable bool) *mcs.MessageErrorResponse {

	var objE mcs.MessageErrorResponse
	objE.Status = s
	objE.Message = m
	objE.Code = code
	objE.OriginalError = originalError
	objE.IsRetryable = isRetryable

	if errorsInfo != nil {
		objE.ErrorsInfo = errorsInfo
	} else {
		objE.ErrorsInfo = make([]mcs.ObjectErrors, 0)
	}

	return &objE
}
