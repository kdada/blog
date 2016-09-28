package services

import (
	"blog/backend/models"

	"github.com/kdada/tinygo/sql"
)

// 用户服务
type UserService struct {
	DB *sql.DB
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
func (this *UserService) Ban(userId int) error {
	return this.DB.Exec("update account set status = 2 where id = $1", userId).Error()
}

// Unban 允许用户登陆
func (this *UserService) Unban(userId int) error {
	return this.DB.Exec("update account set status = 1 where id = $1", userId).Error()
}
