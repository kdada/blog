package models

// 成功返回结果
type SuccessResult struct {
	Code int
	Data interface{}
}

// NewSuccessResult 创建成功返回结果
func NewSuccessResult(result interface{}) *SuccessResult {
	return &SuccessResult{
		0,
		result,
	}
}

// 失败返回结果
type FailureResult struct {
	Code    int
	Message string
}

// NewFailureResult 创建失败返回结果
func NewFailureResult(code int, msg string) *FailureResult {
	return &FailureResult{
		code,
		msg,
	}
}
