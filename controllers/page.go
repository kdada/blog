package controllers

import (
	"blog/services"
	"strconv"

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
	var d, err = layoutViewData(pkg)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	var page, err4 = pkg.ArticleService.Page()
	if err4 != nil {
		context.Processor.Logger.Error(err4.Error())
		return context.NotFound()
	}
	d["ArticlePage"] = page
	if page > 0 {
		if param.Page < 1 {
			param.Page = 1
		}
		var replies, err5 = pkg.ArticleService.Articles(param.Page)
		if err5 != nil || len(replies) <= 0 {
			context.Processor.Logger.Error(err5.Error())
			return context.NotFound()
		}
		d["Articles"] = replies
	}
	d["CurrentPage"] = param.Page
	d["FirstUrl"] = "/"
	d["Url"] = "/p"

	// 计算页码
	var start = param.Page - 3
	if param.Page <= 3 {
		start = 1
	}
	var end = start + 6
	if page-start < 6 {
		end = page
	}

	if start-1 > 6-(end-start) {
		start = end - 6
	} else {
		start = 1
	}
	d["StartPage"] = start
	d["EndPage"] = end
	return context.View("home/index.html", d)
}

// Category 分类列表
func Category(context *web.Context, pkg *ServicePackage, param struct {
	Id   int `!;>0`
	Page int `?;>0`
}) web.GetResult {
	var d, err = layoutViewData(pkg)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	var page, err4 = pkg.ArticleService.CategoryPage(param.Id)
	if err4 != nil || param.Page > page {
		context.Processor.Logger.Error(err4.Error())
		return context.NotFound()
	}
	d["ArticlePage"] = page
	if page > 0 {
		if param.Page < 1 {
			param.Page = 1
		}
		var replies, err5 = pkg.ArticleService.CategoryArticles(param.Id, param.Page)
		if err5 != nil {
			context.Processor.Logger.Error(err5.Error())
			return context.NotFound()
		}
		d["Articles"] = replies
	}
	d["CurrentPage"] = param.Page
	var base = "/c" + strconv.Itoa(param.Id)
	d["FirstUrl"] = base
	d["Url"] = base + "/p"

	// 计算页码
	var start = param.Page - 3
	if param.Page <= 3 {
		start = 1
	}
	var end = start + 6
	if page-start < 6 {
		end = page
	}

	if start-1 > 6-(end-start) {
		start = end - 6
	} else {
		start = 1
	}
	d["StartPage"] = start
	d["EndPage"] = end
	return context.View("home/index.html", d)
}

// Article 文章
func Article(context *web.Context, pkg *ServicePackage, param struct {
	Id int `!;>0`
}) web.GetResult {
	var d, err = layoutViewData(pkg)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	var article, e = pkg.ArticleService.Article(param.Id)
	if e != nil {
		context.Processor.Logger.Error(e.Error())
		return context.NotFound()
	}
	d["Article"] = article
	var pre, err2 = pkg.ArticleService.PreviousArticle(article.Id, article.Category)
	if err2 != nil {
		context.Processor.Logger.Error(err2.Error())
		return context.NotFound()
	}
	d["PreArticle"] = pre
	var next, err3 = pkg.ArticleService.NextArticle(article.Id, article.Category)
	if err3 != nil {
		context.Processor.Logger.Error(err3.Error())
		return context.NotFound()
	}
	d["NextArticle"] = next

	var page, err4 = pkg.ReplyService.Page(article.Id)
	if err4 != nil {
		context.Processor.Logger.Error(err4.Error())
		return context.NotFound()
	}
	d["ReplyPage"] = page
	if page > 0 {
		var replies, err5 = pkg.ReplyService.Replies(article.Id, 1)
		if err5 != nil {
			context.Processor.Logger.Error(err5.Error())
			return context.NotFound()
		}
		d["Replies"] = replies
	}
	return context.View("article/article.html", d)
}

// Manager 管理后台页面
func Manager(context *web.Context, userService *services.UserService) web.GetResult {
	return context.View("manager/index.html")
}
