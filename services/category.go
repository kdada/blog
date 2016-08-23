package services

import (
	"blog/models"

	"github.com/kdada/tinygo/sql"
)

// 分类服务
type CategoryService struct {
	DB *sql.DB
}

// NewCategoryService 创建分类服务
func NewCategoryService() *CategoryService {
	return &CategoryService{
		sql.OpenDefault(),
	}
}

// Categories 获取所有分类
func (this *CategoryService) Categories() ([]*models.Category, error) {
	var v []*models.Category
	var _, err = this.DB.Query("select * from category where status = 1").Scan(&v)
	if err == nil {
		return v, nil
	}
	return nil, err
}
