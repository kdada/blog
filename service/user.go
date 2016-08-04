package service

import (
	"blog/model"
	"errors"
)

// 用户服务
type UserService struct {
}

// Register 注册用户
func (this *UserService) Register(m *model.UserRegister) error {
	if m.Email == "test@test.com" {
		return errors.New("邮箱已存在")
	}
	return nil
}

// Login 登录
func (this *UserService) Login(m *model.UserLogin) (*model.UserInfo, error) {
	if m.Email == "test@test.com" && m.Password == "000000" {
		return &model.UserInfo{1, "Kira", m.Email}, nil
	}
	return nil, errors.New("邮箱或密码错误")
}
