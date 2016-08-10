package services

import (
	"blog/models"
	"errors"
	"fmt"

	"github.com/kdada/tinygo/sql"
)

// 用户服务
type UserService struct {
}

// Register 注册用户
func (this *UserService) Register(m *models.UserRegister) error {
	var db = sql.OpenDefault()
	var id int
	var _, err = db.Query("select id from account where name = $1", m.Name).Scan(&id)
	if err != nil {
		return err
	}
	if id > 0 {
		return errors.New("邮箱已存在")
	}
	_, err = db.Exec("insert into account(email,name,password,salt) values($1,$2,$3,'salt')", m.Email, m.Name, m.Password)
	return err
}

// Login 登录
func (this *UserService) Login(m *models.UserLogin) (*models.UserInfo, error) {
	var u *models.User
	var db = sql.OpenDefault()
	var num, err = db.Query("select * from account where email = $1 and password = $2", m.Email, m.Password).Scan(&u)
	fmt.Println(u, num, err)
	if err == nil && num == 1 {
		return &u.UserInfo, nil
	}
	return nil, errors.New("邮箱或密码错误")
}
