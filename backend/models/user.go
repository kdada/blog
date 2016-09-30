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

// 用户详细信息
type User struct {
	UserDetail
	Password string //密码
	Salt     string //干扰码
}

// 登陆请求参数
type Login struct {
	Email    string `!;/\w[\.\w]*@\w+(\.\w+)+/` //邮箱
	Password string `!;/^\w{6,15}$/`            //密码
}

// 注册请求参数
type Register struct {
	Login
	Name string `!;clen>=2&&clen<=10` //昵称
}
