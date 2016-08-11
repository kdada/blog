package services

import (
	"blog/models"

	"github.com/kdada/tinygo/sql"
)

// 文章服务
type ArticleService struct {
	db *sql.DB
}

// NewArticleService 创建文章服务
func NewArticleService() *ArticleService {
	return &ArticleService{
		sql.OpenDefault(),
	}
}

// NewestArticles 获取最新文章
func (this *ArticleService) NewestArticles() ([]*models.Article, error) {
	var v []*models.Article
	var _, err = this.db.Query("select * from article where status = 1 order by update_time desc limit 10").Scan(&v)
	if err == nil {
		return v, nil
	}
	return nil, err
}
