package controller

import "github.com/kdada/tinygo/web"

type HomeController struct {
	Context *web.Context
}

func (this *HomeController) Index() web.GetResult {
	return this.Context.View("home/index.html", web.ViewData{"SiteName": "测试网站", "Article": "<h1>网页内容测试</h1><h2>网页内容测试2</h2>"})
}
