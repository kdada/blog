package services

import (
	"blog/backend/models"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"

	"github.com/kdada/tinygo/sql"
)

// LinkPassword password连接salt并生成新密码
func LinkPassword(password string, salt string) string {
	var newPwd = salt + password + salt
	var m = md5.New()
	m.Write([]byte(newPwd))
	return hex.EncodeToString(m.Sum(nil))
}

// GeneratePassword 生成密码并返回密码和salt
func GeneratePassword(password string, additional string) (string, string) {
	var salt = strconv.FormatFloat(rand.Float64(), 'e', 20, 64)
	return LinkPassword(password, additional+salt), salt
}

// CheckPassword 检查原密码通过salt后与最终密码是否一致
func CheckPassword(originalPassword string, additional string, salt string, password string) bool {
	return LinkPassword(originalPassword, additional+salt) == password
}

// 用户服务
type UserService struct {
	DB *sql.DB
}

// IsAdmin 判断用户是否是管理员
func (this *UserService) IsAdmin(user int) (bool, error) {
	return user == 1, nil
}

// Check 检查email和password是否合法
func (this *UserService) Check(email string, password string) (*models.UserDetail, error) {
	var user *models.User
	var count, err = this.DB.Query("select * from account where email = $1 and status = 1", email).Scan(&user)
	if err != nil {
		return nil, err
	}
	if count > 0 && CheckPassword(password, email, user.Salt, user.Password) {
		return &user.UserDetail, nil
	}
	return nil, ErrorUnmatchedUser.Error()
}

// Add 添加一个用户
func (this *UserService) Add(email string, name string, password string) error {
	var count int
	var _, err = this.DB.Query("select count(*) from account where email = $1", email).Scan(&count)
	if count > 0 || err != nil {
		return ErrorRepeatedEmail.Error()
	}
	_, err = this.DB.Query("select count(*) from account where name = $1", name).Scan(&count)
	if count > 0 || err != nil {
		return ErrorRepeatedName.Error()
	}
	var pwd, salt = GeneratePassword(password, email)
	return this.DB.Exec("insert into account(email,name,password,salt) values($1,$2,$3,$4)", email, name, pwd, salt).Error()
}

// UserNum 返回用户总数
func (this *UserService) UserNum() (count int, err error) {
	_, err = this.DB.Query("select count(*) from account").Scan(&count)
	return
}

// ListAll 列出指定页码的用户详细信息
func (this *UserService) ListAll(page int, count int) (details []*models.UserDetail, err error) {
	_, err = this.DB.Query("select * from account order by status asc,id asc limit $1 offset $2", count, (page-1)*count).Scan(&details)
	return
}

// Ban 禁止用户登陆
func (this *UserService) Ban(user int, reason string) error {
	return this.DB.Exec("update account set status = 2,reason=$2 where id = $1", user, reason).Error()
}

// Unban 允许用户登陆
func (this *UserService) Unban(user int) error {
	return this.DB.Exec("update account set status = 1 where id = $1", user).Error()
}
