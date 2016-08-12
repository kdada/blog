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
func (this *ArticleService) NewestArticles() ([]*models.ArticleSummary, error) {
	var v []*models.ArticleSummary
	var _, err = this.db.Query("select * from article where status = 1 order by update_time desc limit 10").Scan(&v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

// Page 返回所有文章总页数
func (this *ArticleService) Page() (int, error) {
	var v int
	var _, err = this.db.Query(`select count(*) from article where status = 1`).Scan(&v)
	if err == nil {
		var page = v / 10
		if v%10 > 0 {
			page++
		}
		return page, nil
	}
	return 0, err
}

// Articles 返回指定页数的文章
func (this *ArticleService) Articles(page int) ([]*models.ArticleSummary, error) {
	var v []*models.ArticleSummary
	var _, err = this.db.Query(`
select * 
from article 
where status = 1 
order by top,create_time desc 
offset $1
limit $2`, (page-1)*10, 10).Scan(&v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

// CategoryPage 返回所有文章总页数
func (this *ArticleService) CategoryPage(category int) (int, error) {
	var v int
	var _, err = this.db.Query(`select count(*) from article where category = $1 and status = 1`, category).Scan(&v)
	if err == nil {
		var page = v / 10
		if v%10 > 0 {
			page++
		}
		return page, nil
	}
	return 0, err
}

// CategoryArticles 返回指定页数的文章
func (this *ArticleService) CategoryArticles(category int, page int) ([]*models.ArticleSummary, error) {
	var v []*models.ArticleSummary
	var _, err = this.db.Query(`
select * 
from article 
where status = 1 and category = $1 
order by top,create_time desc 
offset $2
limit $3`, category, (page-1)*10, 10).Scan(&v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

// Article 获取指定id的文章
func (this *ArticleService) Article(id int) (*models.Article, error) {
	var v *models.Article
	var _, err = this.db.Query(`
select a.*,c.name
from category c,
(select * 
from article 
where status = 1 and id = $1) a
where c.id = a.category`, id).Scan(&v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

// PreviousArticle 指定id的上一篇文章
func (this *ArticleService) PreviousArticle(id, category int) (*models.Article, error) {
	var v *models.Article
	var count, err = this.db.Query(`
select * 
from article 
where status = 1 and id < $1 and category = $2
order by id desc
limit 1
`, id, category).Scan(&v)
	if err != nil {
		return nil, err
	}
	if count <= 0 {
		return nil, nil
	}
	return v, nil
}

// NextArticle 指定id的下一篇文章
func (this *ArticleService) NextArticle(id, category int) (*models.Article, error) {
	var v *models.Article
	var count, err = this.db.Query(`
select * 
from article 
where status = 1 and id > $1 and category = $2
order by id asc
limit 1
`, id, category).Scan(&v)
	if err != nil {
		return nil, err
	}
	if count <= 0 {
		return nil, nil
	}
	return v, nil
}
