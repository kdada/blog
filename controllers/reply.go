package controllers

import (
	"blog/models"
	"blog/services"

	"github.com/kdada/tinygo/web"
)

// 评论控制器
type ReplyController struct {
	Context      *web.Context
	ReplyService *services.ReplyService
}

// Page 返回指定页码的评论
func (this *ReplyController) List(param struct {
	Article int `!;>0`
	Page    int `!;>0`
}) web.PostResult {
	var r, err = this.ReplyService.Replies(param.Article, param.Page)
	if err != nil {
		this.Context.Api(models.NewErrorResult(err))
	}
	return this.Context.Api(models.NewSuccessResult(r))
}

// New 创建评论
func (this *ReplyController) New(param struct {
	Article int    `!;>0`
	Reply   int    `!;>=0`
	Content string `!;len>=2`
}) web.PostResult {
	var login, ok = this.Context.Session.Bool("Login")
	if ok && login {
		var userInfo, _ = this.Context.Session.Value("User")
		var err = this.ReplyService.New(param.Article, param.Reply, userInfo.(*models.UserInfo).Id, param.Content)
		if err != nil {
			return this.Context.Api(models.NewErrorResult(err))
		}
		return this.Context.Api(models.NewSuccessResult(nil))
	}
	return this.Context.NotFound()
}
