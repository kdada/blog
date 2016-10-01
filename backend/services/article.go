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

// NewestArticles 获取最新文章
func (this *ArticleService) NewestArticles() (details []*models.ArticleDetail, err error) {
	_, err = this.DB.Query("select * from article where status = 1 order by id desc limit 10").Scan(&details)
	return
}

// AvailableArticle 获取指定id的正常状态的文章详情(html作为content)
func (this *ArticleService) AvailableArticle(article int) (detail *models.ArticleDetail, err error) {
	_, err = this.DB.Query("select a.id,a.title,a.html as content,a.category,a.top,a.create_time,a.update_time,a.status,c.name from article a,category c where a.id = $1 and a.status = 1 and  a.category = c.id", article).Scan(&detail)
	return
}

// AvailableNum 返回正常状态的文章总数
func (this *ArticleService) AvailableNum() (count int, err error) {
	_, err = this.DB.Query("select count(*) from article where status = 1").Scan(&count)
	return
}

// ListAvailable 列出指定页码的处于正常状态的文章(summary作为content)
func (this *ArticleService) ListAvailable(page int, count int) (details []*models.ArticleDetail, err error) {
	_, err = this.DB.Query(`select a.id,a.title,a.summary as content,a.category,c.name,
	a.top,a.create_time,a.update_time,a.status
	from blog.article a,blog.category c
	where a.status = 1 and a.category = c.id
	order by a.top desc,a.id desc
	limit $1 offset $2`, count, (page-1)*count).Scan(&details)
	return
}

// Article 获取指定id的文章详情
func (this *ArticleService) Article(article int) (detail *models.ArticleDetail, err error) {
	_, err = this.DB.Query("select a.*,c.name from article a,category c where a.id = $1 and a.category = c.id", article).Scan(&detail)
	return
}

// ArticleNum 返回正常状态和隐藏状态的文章总数
func (this *ArticleService) ArticleNum() (count int, err error) {
	_, err = this.DB.Query("select count(*) from article where status = 1 or status = 2").Scan(&count)
	return
}

// ListAll 列出指定页码的处于正常状态和隐藏状态的文章(summary作为content)
func (this *ArticleService) ListAll(page int, count int) (details []*models.ArticleDetail, err error) {
	_, err = this.DB.Query(`select a.id,a.title,a.summary as content,a.category,c.name,
	a.top,a.create_time,a.update_time,a.status
	from article a,category c
	where (a.status = 1 or a.status = 2) and a.category = c.id
	order by a.top desc,a.status asc,a.id desc
	limit $1 offset $2`, count, (page-1)*count).Scan(&details)
	return
}

// SpecArticleNum 返回正常状态和隐藏状态的文章总数
func (this *ArticleService) SpecArticleNum(category int) (count int, err error) {
	_, err = this.DB.Query("select count(*) from article where category = $1 and (status = 1 or status = 2)", category).Scan(&count)
	return
}

// SpecAvailableNum 返回正常状态的文章总数
func (this *ArticleService) SpecAvailableNum(category int) (count int, err error) {
	_, err = this.DB.Query("select count(*) from article where category = $1 and status = 1", category).Scan(&count)
	return
}

// ListSpecAvailable 列出指定页码的处于正常状态的文章(summary作为content)
func (this *ArticleService) ListSpecAvailable(category int, page int, count int) (details []*models.ArticleDetail, err error) {
	_, err = this.DB.Query(`select a.id,a.title,a.summary as content,a.category,c.name,
	a.top,a.create_time,a.update_time,a.status
	from blog.article a,blog.category c
	where category = $3 and a.status = 1 and a.category = c.id
	order by a.top desc,a.status asc,a.id desc
	limit $1 offset $2`, count, (page-1)*count, category).Scan(&details)
	return
}

// ListSpec 列出指定页码的处于正常状态和隐藏状态的文章(summary作为content)
func (this *ArticleService) ListSpec(category int, page int, count int) (details []*models.ArticleDetail, err error) {
	_, err = this.DB.Query(`select a.id,a.title,a.summary as content,a.category,c.name,
	a.top,a.create_time,a.update_time,a.status
	from blog.article a,blog.category c
	where category = $3 and (a.status = 1 or a.status = 2) and a.category = c.id
	order by a.top desc,a.status asc,a.id desc
	limit $1 offset $2`, count, (page-1)*count, category).Scan(&details)
	return
}

// Create 创建一篇文章并返回文章id
func (this *ArticleService) Create(article *models.ArticleInfo) (id int, err error) {
	_, err = this.DB.Query("insert into article(category,title,content,summary,html) values($1,$2,$3,$4,$5) returning id", article.Category, article.Title, article.Content, article.Summary, article.Html).Scan(&id)
	return
}

// Update 更新一篇文章
func (this *ArticleService) Update(article *models.ArticleData) error {
	return this.DB.Exec("update article set category=$1,title=$2,content=$3,summary=$4,html=$5,update_time=now() where id = $6", article.Category, article.Title, article.Content, article.Summary, article.Html, article.Id).Error()
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

// PreviousArticle 指定文章同一分类的上一篇文章
func (this *ArticleService) PreviousArticle(article, category int) (*models.ArticleDetail, error) {
	var detail *models.ArticleDetail
	var count, err = this.DB.Query("select * from article where status = 1 and id < $1 and category = $2 order by id desc limit 1", article, category).Scan(&detail)
	if err == nil && count <= 0 {
		return nil, nil
	}
	return detail, err
}

// NextArticle 指定文章同一分类的下一篇文章
func (this *ArticleService) NextArticle(article, category int) (*models.ArticleDetail, error) {
	var detail *models.ArticleDetail
	var count, err = this.DB.Query("select * from article where status = 1 and id > $1 and category = $2 order by id asc limit 1", article, category).Scan(&detail)
	if err == nil && count <= 0 {
		return nil, nil
	}
	return detail, err
}
