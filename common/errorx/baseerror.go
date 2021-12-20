package errorx

const defaultCode = 400

type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CodeErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCodeError(code int, message string) error {
	return &CodeError{
		Code:    code,
		Message: message,
	}
}

func NewDefaultError(message string) error {
	return NewCodeError(defaultCode, message)
}

func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
	}
}
