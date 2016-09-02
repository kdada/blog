package services

import (
	"blog/models"

	"github.com/kdada/tinygo/sql"
)

// 用户服务
type UserService struct {
	DB *sql.DB
}

// NewUserService 创建用户服务
func NewUserService() *UserService {
	return &UserService{
		sql.OpenDefault(),
	}
}

// IsAdmin 返回指定id的用户是否是管理员
func (this *UserService) IsAdmin(id int) bool {
	return id == 1
}

// Register 注册用户
func (this *UserService) Register(email, name, password string) error {
	var id int
	var _, err = this.DB.Query("select id from account where email = $1", email).Scan(&id)
	if err != nil {
		return err
	}
	if id > 0 {
		return models.ErrorExistentEmail.Error()
	}
	_, err = this.DB.Query("select id from account where name = $1", name).Scan(&id)
	if err != nil {
		return err
	}
	if id > 0 {
		return models.ErrorExistentName.Error()
	}
	return this.DB.Exec("insert into account(email,name,password,salt) values($1,$2,$3,'salt')", email, name, password).Error()
}

// Login 登录
func (this *UserService) Login(email, password string) (*models.UserInfo, error) {
	var u *models.User
	var num, err = this.DB.Query("select * from account where email = $1 and password = $2", email, password).Scan(&u)
	if err != nil {
		return nil, err
	}
	if num == 1 {
		return &u.UserInfo, nil
	}
	return nil, models.ErrorInvalidLogin.Error()
}
