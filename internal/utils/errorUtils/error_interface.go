package errorUtils

type Error struct {
	Code                 int    `json:"code"`
	Status               int    `json:"status,omitempty"`
	ErrorCode            string `json:"error_code,omitempty"`
	Message              string `json:"message"`
	InternalErrorMessage string `json:"-"`
}

func (e *Error) Error() string {
	if e.InternalErrorMessage != "" {
		return e.InternalErrorMessage
	}
	return e.Message
}

func (e *Error) WithInternalError(inErr error) *Error {
	e.InternalErrorMessage = inErr.Error()
	return e
}
