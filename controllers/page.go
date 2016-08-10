package controllers

import (
	"blog/services"

	"github.com/kdada/tinygo/web"
)

// layoutViewData 返回布局用的数据
func layoutViewData(articleService *services.ArticleService) (web.ViewData, error) {
	var c, err = articleService.Categories()
	if err != nil {
		return nil, err
	}
	var a, err1 = articleService.NewestArticles()
	if err1 != nil {
		return nil, err1
	}
	var r, err2 = articleService.NewestReplies()
	if err2 != nil {
		return nil, err2
	}
	return web.ViewData{"Categories": c, "NewestArticles": a, "NewestReplies": r}, nil
}

// Index 首页
func Index(context *web.Context, articleService *services.ArticleService) web.GetResult {
	var d, err = layoutViewData(articleService)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	return context.View("home/index.html", d)
}

// Article 文章
func Article(context *web.Context, articleService *services.ArticleService, param struct {
	Id int `!;>0`
}) web.GetResult {
	context.Processor.Logger.Debug(param)
	var d, err = layoutViewData(articleService)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	return context.View("article/article.html", d)
}

// Category 分类列表
func Category(context *web.Context, articleService *services.ArticleService, param struct {
	Id   int `!;>0`
	Page int `?;>=0`
}) web.GetResult {
	context.Processor.Logger.Debug(param)
	var d, err = layoutViewData(articleService)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	return context.View("home/index.html", d)
}

// Manager 管理后台页面
func Manager(context *web.Context) web.GetResult {
	return context.View("manager/index.html")
}
