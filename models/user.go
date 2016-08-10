package models

import (
	"time"
)

// 用户登录模型
type UserLogin struct {
	Email    string `!;/\w+[\.\w]*@\w+(\.\w+)+/` //邮箱
	Password string `!;len>=6&&len<=15`          //密码
}

// 用户注册模型
type UserRegister struct {
	UserLogin
	Name string `!;len>=2&&len<=10` //昵称
}

// 用户信息
type UserInfo struct {
	Id    int    //用户id
	Name  string //用户昵称
	Email string //用户邮箱
}

// 用户记录
type User struct {
	UserInfo
	Password   string    //用户密码
	CreateTime time.Time //注册时间
	Status     int       //状态
	Reason     string    //状态原因
}
