package models

import (
	"regexp"
	"strconv"
)

// 成功返回结果,Code为0
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

// Error格式的正则
var regErr = regexp.MustCompile(`^[a-zA-Z\d]+?\(([A-Z])(\d+)\):(.+?)$`)

// NewErrorResult 根据err返回结果,err需要符合Error的错误格式
func NewErrorResult(err error) *FailureResult {
	if err != nil {
		var v = regErr.FindStringSubmatch(err.Error())
		if len(v) == 4 && v[1] == "P" {
			//格式符合要求,并且错误码以P开头
			var codeStr = v[2]
			var code, _ = strconv.Atoi(codeStr)
			var prefix = int(v[1][0])
			for i := 0; i < len(codeStr); i++ {
				prefix *= 10
			}
			code += prefix
			return &FailureResult{code, v[3]}
		}
	}
	return FailureInternal
}

// 应该使用error生成FailureResult,如果err不符合格式要求,则应该返回预定义的FailureResult
var (
	FailureInternal = &FailureResult{1000, "网站出现了一些问题,暂时不能提供服务"}
)
