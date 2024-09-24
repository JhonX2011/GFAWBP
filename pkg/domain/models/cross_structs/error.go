package crossstructs

// MessageErrorResponse default structure from basic error
type MessageErrorResponse struct {
	Status        string         `json:"status,omitempty"`
	Message       string         `json:"message"`
	ErrorsInfo    []ObjectErrors `json:"errorsInfo,omitempty"`
	Code          int            `json:"code"`
	OriginalError error
	IsRetryable   bool
}

// ObjectErrors Custom error handling information.
type ObjectErrors struct {
	// detail of error
	DetailError string   `json:"detail_error,omitempty"`
	StackTrace  []string `json:"stack_trace,omitempty"`
}

// Error get message error.
func (e *MessageErrorResponse) Error() string {
	return e.Message
}
