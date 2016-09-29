package controllers

import (
	"blog/backend/services"

	"github.com/kdada/tinygo/web"
)

// 服务打包
type ServicePackage struct {
	ArticleService  *services.ArticleService  //文章服务
	CategoryService *services.CategoryService // 分类服务
	ReplyService    *services.ReplyService    // 评论服务
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
func Index(context *web.Context, pkg *ServicePackage, param struct {
	Page int `?;>0`
}) web.GetResult {
	return context.View("pages/index.html")
}

// Manager 返回管理页面
func Manager(context *web.Context) web.GetResult {
	return context.View("manager/index.html")
}
