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
