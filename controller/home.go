package controller

import (
	"blog/model"
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

func (this *HomeController) Article() web.GetResult {
	return this.Context.View("article/article.html")
}

func (this *HomeController) Login(m *model.UserLogin) web.PostResult {
	var s = new(service.UserService)
	var userInfo, err = s.Login(m)
	if err != nil {
		return this.Context.Api(model.NewFailureResult(1, err.Error()))
	}
	this.Context.Session.SetBool("Login", true)
	this.Context.Session.SetValue("User", userInfo)
	return this.Context.Api(model.NewSuccessResult(userInfo))
}
