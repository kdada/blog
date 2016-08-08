package controller

import (
	"blog/service"

	"github.com/kdada/tinygo/web"
)

type HomeController struct {
	Context        *web.Context
	ArticleService *service.ArticleService
}

func (this *HomeController) Index() web.GetResult {
	var c, _ = this.ArticleService.Categories()
	var a, _ = this.ArticleService.NewestArticles()
	var r, _ = this.ArticleService.NewestReplies()
	return this.Context.View("home/index.html", web.ViewData{"Categories": c, "NewestArticles": a, "NewestReplies": r})
}

func (this *HomeController) Article(ctx *web.Context) web.GetResult {
	return ctx.View("article/article.html")
}
