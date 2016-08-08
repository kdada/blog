package services

import (
	"blog/models"
	"errors"
)

// 用户服务
type UserService struct {
}

var users = map[string]*models.User{}

// Register 注册用户
func (this *UserService) Register(m *models.UserRegister) error {
	var _, ok = users[m.Email]
	if ok {
		return errors.New("邮箱已存在")
	}
	users[m.Email] = &models.User{models.UserInfo{len(users) + 1, m.Name, m.Email}, m.Password}
	return nil
}

// Login 登录
func (this *UserService) Login(m *models.UserLogin) (*models.UserInfo, error) {
	var u, ok = users[m.Email]
	if ok && m.Password == u.Password {
		return &u.UserInfo, nil
	}
	return nil, errors.New("邮箱或密码错误")
}
