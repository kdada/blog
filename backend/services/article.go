package services

import (
	"blog/backend/models"
	"time"

	"github.com/kdada/tinygo/sql"
)

// 文章服务
type ArticleService struct {
	DB *sql.DB
}

// AvailableArticle 获取指定id的正常状态的文章详情
func (this *ArticleService) AvailableArticle(article int) (detail *models.ArticleDetail, err error) {
	_, err = this.DB.Query("select * from article where id = $1 and status = 1", article).Scan(&detail)
	return
}

// AvailableNum 返回正常状态的文章总数
func (this *ArticleService) AvailableNum() (count int, err error) {
	_, err = this.DB.Query("select count(*) from article where status = 1").Scan(&count)
	return
}

// ListAvailable 列出指定页码的处于正常状态的文章
func (this *ArticleService) ListAvailable(page int, count int) (details []*models.ArticleDetail, err error) {
	_, err = this.DB.Query("select * from article where status = 1 order by top desc,id asc limit $1 offset $2", count, (page-1)*count).Scan(&details)
	return
}

// Article 获取指定id的文章详情
func (this *ArticleService) Article(article int) (detail *models.ArticleDetail, err error) {
	_, err = this.DB.Query("select * from article where id = $1", article).Scan(&detail)
	return
}

// ArticleNum 返回正常状态和隐藏状态的文章总数
func (this *ArticleService) ArticleNum() (count int, err error) {
	_, err = this.DB.Query("select count(*) from article where status = 1 or status = 2").Scan(&count)
	return
}

// ListAll 列出指定页码的处于正常状态和隐藏状态的文章
func (this *ArticleService) ListAll(page int, count int) (details []*models.ArticleDetail, err error) {
	_, err = this.DB.Query("select * from article where status = 1 or status = 2 order by top desc,status asc,id asc limit $1 offset $2", count, (page-1)*count).Scan(&details)
	return
}

// SpecArticleNum 返回正常状态和隐藏状态的文章总数
func (this *ArticleService) SpecArticleNum(category int) (count int, err error) {
	_, err = this.DB.Query("select count(*) from article where cateory = $1 and status = 1 or status = 2", category).Scan(&count)
	return
}

// ListSpec 列出指定页码的处于正常状态和隐藏状态的文章
func (this *ArticleService) ListSpec(category int, page int, count int) (details []*models.ArticleDetail, err error) {
	_, err = this.DB.Query("select * from article where category = $3 and status = 1 or status = 2 order by top desc,status asc,id asc  limit $1 offset $2", count, (page-1)*count, category).Scan(&details)
	return
}

// Create 创建一篇文章并返回文章id
func (this *ArticleService) Create(article *models.ArticleInfo) (int, error) {
	return this.DB.Exec("insert into article(category,title,content) values($1,$2,$3)", article.Category, article.Title, article.Content).LastInsertId()
}

// Update 更新一篇文章
func (this *ArticleService) Update(article *models.ArticleData) error {
	return this.DB.Exec("update article set title=$1,content=$2,update_time=now() where id = $3", article.Title, article.Content, article.Id).Error()
}

// Move 移动到新的分类
func (this *ArticleService) Move(article int, category int) error {
	return this.DB.Exec("update article set category=$1 where id = $2", category, article).Error()
}

// Hide 隐藏文章
func (this *ArticleService) Hide(article int) error {
	return this.DB.Exec("update article set status=2 where id = $1", article).Error()
}

// Show 显示文章
func (this *ArticleService) Show(article int) error {
	return this.DB.Exec("update article set status=1 where id = $1", article).Error()
}

// Top 置顶文章
func (this *ArticleService) Top(article int) error {
	return this.DB.Exec("update article set top=$1 where id = $2", time.Now().Unix(), article).Error()
}

// Untop 取消置顶
func (this *ArticleService) Untop(article int) error {
	return this.DB.Exec("update article set top=0 where id = $1", article).Error()
}

// Delete 删除文章
func (this *ArticleService) Delete(article int) error {
	return this.DB.Exec("update article set status=3 where id = $1", article).Error()
}
