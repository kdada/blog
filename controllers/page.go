package controllers

import (
	"blog/services"

	"github.com/kdada/tinygo/web"
)

// 服务打包
type ServicePackage struct {
	ArticleService  *services.ArticleService  //文章服务
	CategoryService *services.CategoryService // 分类服务
	ReplyService    *services.ReplyService    // 回复服务
}

// layoutViewData 返回布局用的数据
func layoutViewData(pkg *ServicePackage) (web.ViewData, error) {
	var c, err = pkg.CategoryService.Categories()
	if err != nil {
		return nil, err
	}
	var a, err1 = pkg.ArticleService.NewestArticles()
	if err1 != nil {
		return nil, err1
	}
	var r, err2 = pkg.ReplyService.NewestReplies()
	if err2 != nil {
		return nil, err2
	}
	return web.ViewData{"Categories": c, "NewestArticles": a, "NewestReplies": r}, nil
}

// Index 首页
func Index(context *web.Context, pkg *ServicePackage) web.GetResult {
	var d, err = layoutViewData(pkg)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	return context.View("home/index.html", d)
}

// Article 文章
func Article(context *web.Context, pkg *ServicePackage, param struct {
	Id int `!;>0`
}) web.GetResult {
	context.Processor.Logger.Debug(param)
	var d, err = layoutViewData(pkg)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	return context.View("article/article.html", d)
}

// Category 分类列表
func Category(context *web.Context, pkg *ServicePackage, param struct {
	Id   int `!;>0`
	Page int `?;>=0`
}) web.GetResult {
	context.Processor.Logger.Debug(param)
	var d, err = layoutViewData(pkg)
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
