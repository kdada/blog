package controllers

import (
	"blog/backend/models"
	"blog/backend/services"
	"encoding/json"
	"io/ioutil"

	"github.com/kdada/tinygo/web"
)

// 文章控制器
type ArticleController struct {
	BaseController
	ArticleService *services.ArticleService
}

// List 显示文章列表,每页10行
func (this *ArticleController) List(params struct {
	Category int //分类 为0时视为所有分类
	Page     int `!;>0` //页码
}) web.PostResult {
	var r []*models.ArticleDetail
	var err error
	if params.Category == 0 {
		r, err = this.ArticleService.ListAll(params.Page, pageCount)
	} else {
		r, err = this.ArticleService.ListSpec(params.Category, params.Page, pageCount)
	}
	return this.returnPostResult(r, err)
}

// ListNum 返回页数
func (this *ArticleController) ListNum(params struct {
	Category int //分类 为0时视为所有分类
}) web.PostResult {
	var count = 0
	var err error
	if params.Category == 0 {
		count, err = this.ArticleService.ArticleNum()
	} else {
		count, err = this.ArticleService.SpecArticleNum(params.Category)
	}
	var result interface{} = nil
	if err == nil {
		result = map[string]int{
			"Count": calculatePages(count),
		}
	}
	return this.returnPostResult(result, err)

}

// Article 返回指定文章信息
func (this *ArticleController) Article(params struct {
	Article int `!;>0` //文章id
}) web.PostResult {
	return this.returnPostResult(this.ArticleService.Article(params.Article))
}

// jsonRequest 处理json格式的请求
func (this *ArticleController) jsonRequest(obj interface{}) error {
	var all, err = ioutil.ReadAll(this.Context.HttpContext.Request.Body)
	if err == nil {
		return json.Unmarshal(all, obj)
	}
	return err
}

// Create 创建一篇文章并返回文章id
func (this *ArticleController) Create() web.PostResult {
	var article = new(models.ArticleInfo)
	var err = this.jsonRequest(article)
	var result interface{}
	if err == nil {
		if article.Validate() {
			var id = 0
			id, err = this.ArticleService.Create(article)
			if err == nil {
				result = map[string]int{
					"Id": id,
				}
				return this.Context.Json(models.NewSuccessResult(result))
			}
		}
	}
	return this.Context.Json(models.NewErrorResult(err))
}

// Update 更新一篇文章
func (this *ArticleController) Update() web.PostResult {
	var article = new(models.ArticleData)
	var err = this.jsonRequest(article)
	if err == nil && article.Validate() {
		return this.returnPostResult(nil, this.ArticleService.Update(article))
	}
	return this.Context.Json(models.NewErrorResult(err))
}

// Move 移动到新的分类
func (this *ArticleController) Move(params struct {
	Article  int // 文章id
	Category int // 文章分类id
}) web.PostResult {
	return this.returnPostResult(nil, this.ArticleService.Move(params.Article, params.Category))
}

// Hide 隐藏文章
func (this *ArticleController) Hide(params struct {
	Article int // 文章id
}) web.PostResult {
	return this.returnPostResult(nil, this.ArticleService.Hide(params.Article))
}

// Show 显示文章
func (this *ArticleController) Show(params struct {
	Article int // 文章id
}) web.PostResult {
	return this.returnPostResult(nil, this.ArticleService.Show(params.Article))
}

// Top 置顶文章
func (this *ArticleController) Top(params struct {
	Article int // 文章id
}) web.PostResult {
	return this.returnPostResult(nil, this.ArticleService.Top(params.Article))
}

// Untop 取消置顶
func (this *ArticleController) Untop(params struct {
	Article int // 文章id
}) web.PostResult {
	return this.returnPostResult(nil, this.ArticleService.Untop(params.Article))
}

// Delete 删除文章
func (this *ArticleController) Delete(params struct {
	Article int // 文章id
}) web.PostResult {
	return this.returnPostResult(nil, this.ArticleService.Show(params.Article))
}
