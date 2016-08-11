package models

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
//  格式:错误名称(错误编号前缀+错误编号):错误描述
//  1.错误名称应当具有一定的描述性,最好只包含英文和数字
//  2.错误编号前缀通常为一个大写字母,注意与tinygo框架中已经使用的前缀区分(已使用前缀:CNMLRSQTUVW),P表示Project
//  3.错误编号一般为5位
const (
	ErrorExistentEmail Error = "ErrorExistentEmail(P10000):邮箱已存在"
	ErrorExistentName  Error = "ErrorExistentName(P10001):昵称已存在"
	ErrorInvalidLogin  Error = "ErrorInvalidLogin(P10010):邮箱或密码错误"
)
