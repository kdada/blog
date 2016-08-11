package services

import (
	"blog/models"

	"github.com/kdada/tinygo/sql"
)

// 用户服务
type UserService struct {
	db *sql.DB
}

// NewUserService 创建用户服务
func NewUserService() *UserService {
	return &UserService{
		sql.OpenDefault(),
	}
}

// Register 注册用户
func (this *UserService) Register(email, name, password string) error {
	var id int
	var _, err = this.db.Query("select id from account where email = $1", email).Scan(&id)
	if err != nil {
		return err
	}
	if id > 0 {
		return models.ErrorExistentEmail.Error()
	}
	_, err = this.db.Query("select id from account where name = $1", name).Scan(&id)
	if err != nil {
		return err
	}
	if id > 0 {
		return models.ErrorExistentName.Error()
	}
	_, err = this.db.Exec("insert into account(email,name,password,salt) values($1,$2,$3,'salt')", email, name, password)
	return err
}

// Login 登录
func (this *UserService) Login(email, password string) (*models.UserInfo, error) {
	var u *models.User
	var num, err = this.db.Query("select * from account where email = $1 and password = $2", email, password).Scan(&u)
	if err != nil {
		return nil, err
	}
	if num == 1 {
		return &u.UserInfo, nil
	}
	return nil, models.ErrorInvalidLogin.Error()
}
