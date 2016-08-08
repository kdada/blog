package models

type SuccessResult struct {
	Code int
	Data interface{}
}

func NewSuccessResult(result interface{}) *SuccessResult {
	return &SuccessResult{
		0,
		result,
	}
}

type FailureResult struct {
	Code    int
	Message string
}

func NewFailureResult(code int, msg string) *FailureResult {
	return &FailureResult{
		code,
		msg,
	}
}
