package models

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
