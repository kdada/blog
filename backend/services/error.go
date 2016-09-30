package services

import (
	"errors"
	"fmt"
)

// 错误信息
type Error string

// Format 格式化错误信息并生成新的错误信息
func (this Error) Format(data ...interface{}) Error {
	return Error(fmt.Sprintf(string(this), data...))
}

// Error 生成error类型
func (this Error) Error() error {
	return errors.New(string(this))
}

// String 返回错误字符串描述
func (this Error) String() string {
	return string(this)
}

// 错误码
const (
	ErrorRepeatedEmail Error = "ErrorRepeatedEmail(P100000):邮箱已存在"
	ErrorRepeatedName  Error = "ErrorRepeatedName(P100001):昵称已存在"
	ErrorUnmatchedUser Error = "ErrorUnmatchedUser(P100002):邮箱或密码不正确"
)
