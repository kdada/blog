package models

import "time"

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
