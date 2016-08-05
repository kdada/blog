package controller

import (
	"blog/service"

	"github.com/kdada/tinygo/web"
)

type HomeController struct {
	Context *web.Context
}

func (this *HomeController) Index() web.GetResult {
	var s = new(service.ArticleService)
	var c, _ = s.Categories()
	var a, _ = s.NewestArticles()
	var r, _ = s.NewestReplies()
	return this.Context.View("home/index.html", web.ViewData{"Categories": c, "NewestArticles": a, "NewestReplies": r})
}

func (this *HomeController) Article(Context *web.Context) web.GetResult {
	return this.Context.View("article/article.html")
}
