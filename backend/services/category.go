package services

import (
	"blog/backend/models"

	"github.com/kdada/tinygo/sql"
)

// 分类服务
type CategoryService struct {
	DB *sql.DB
}

// AvailableNum 返回可见分类总数
func (this *CategoryService) AvailableNum() (count int, err error) {
	_, err = this.DB.Query("select count(*) from category where status = 1").Scan(&count)
	return
}

// ListAvailable 列出指定页码的处于正常状态的用户详细信息
func (this *CategoryService) ListAvailable(page int, count int) (details []*models.CategoryDetail, err error) {
	_, err = this.DB.Query("select * from category where status = 1 order by id asc limit $1 offset $2", count, (page-1)*count).Scan(&details)
	return
}

// CategoryNum 返回分类总数
func (this *CategoryService) CategoryNum() (count int, err error) {
	_, err = this.DB.Query("select count(*) from category where status = 1 or status = 2").Scan(&count)
	return
}

// ListAll 列出指定页码的处于正常状态和隐藏状态的用户详细信息
func (this *CategoryService) ListAll(page int, count int) (details []*models.CategoryDetail, err error) {
	_, err = this.DB.Query("select * from category where status = 1 or status = 2 order by status asc,id asc limit $1 offset $2", count, (page-1)*count).Scan(&details)
	return
}

// Hide 隐藏分类
func (this *CategoryService) Hide(categoryId int) error {
	return this.DB.Exec("update category set status = 2 where id = $1", categoryId).Error()
}

// Show 显示分类
func (this *CategoryService) Show(categoryId int) error {
	return this.DB.Exec("update category set status = 1 where id = $1", categoryId).Error()
}

// Delete 删除分类,同时删除分类下所有文章
func (this *CategoryService) Delete(categoryId int) error {
	var err = this.DB.Begin()
	if err != nil {
		return err
	}
	err = this.DB.Exec("update category set status = 3 where id = $1", categoryId).Error()
	if err != nil {
		this.DB.Rollback()
		return err
	}
	err = this.DB.Exec("update article set status = 3 where category = $1", categoryId).Error()
	if err != nil {
		this.DB.Rollback()
		return err
	}
	err = this.DB.Commit()
	if err != nil {
		return err
	}
	return nil
}

// Create 创建分类
func (this *CategoryService) Create(name string) error {
	return this.DB.Exec("insert into category(name) values($1)", name).Error()
}
