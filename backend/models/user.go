package models

import (
	"time"
)

// 用户详细信息
type UserDetail struct {
	Id         int       //用户id
	Email      string    //用户邮箱
	Name       string    //用户昵称
	CreateTime time.Time //创建时间
	Status     int       //状态码:1-正常,2-禁止登录
	Reason     string    //处于当前状态的原因
}
