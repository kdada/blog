package controllers

import (
	"blog/backend/models"
	"blog/backend/services"
	"fmt"

	"github.com/kdada/tinygo/web"
)

// 回复控制器
type ReplyController struct {
	BaseController
	UserSerivce    *services.UserService
	ArticleService *services.ArticleService
	ReplyService   *services.ReplyService
}

// Create 创建回复
func (this *ReplyController) Create(params struct {
	Article int    `!;>0`                  //文章id
	Reply   int    `!;>=0`                 //回复对象id
	Content string `!;clen>=2&&clen<=1000` //回复内容
}) web.PostResult {
	//检查用户是否已经登陆
	var obj, ok = this.Context.Session.Value("User")
	if ok {
		var user = obj.(*models.UserDetail)
		var article, err = this.ArticleService.AvailableArticle(params.Article)
		fmt.Println(article, err)
		if err == nil && article != nil && article.Id > 0 {
			return this.returnPostResult(nil, this.ReplyService.Create(params.Article, params.Reply, user.Id, params.Content))
		}
		//非法请求,禁用用户
		this.UserSerivce.Ban(user.Id, fmt.Sprint("非法请求:", params))
	}
	return this.Context.NotFound()
}

// List 显示回复列表,每页10行
func (this *ReplyController) List(params struct {
	Article int `!;>0` //文章id
	Page    int `!;>0` //页码
}) web.PostResult {
	return this.returnPostResult(this.ReplyService.ListAll(params.Article, params.Page, pageCount))
}

// ListNum 返回页数
func (this *ReplyController) ListNum(params struct {
	Article int `!;>0` //文章id
}) web.PostResult {
	var count, err = this.ReplyService.ReplyNum(params.Article)
	var result interface{} = nil
	if err == nil {
		result = map[string]int{
			"Count": calculatePages(count),
		}
	}
	return this.returnPostResult(result, err)
}
