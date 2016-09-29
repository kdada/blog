package controllers

import (
	"blog/backend/models"
	"blog/backend/services"
	"strconv"

	"github.com/kdada/tinygo/web"
)

// 服务打包
type ServicePackage struct {
	ArticleService  *services.ArticleService  //文章服务
	CategoryService *services.CategoryService // 分类服务
	ReplyService    *services.ReplyService    // 评论服务
}

// layoutViewData 返回边栏数据
func layoutViewData(pkg *ServicePackage) (web.ViewData, error) {
	var c, err = pkg.CategoryService.AvailableCategories()
	if err != nil {
		return nil, err
	}
	a, err := pkg.ArticleService.NewestArticles()
	if err != nil {
		return nil, err
	}
	r, err := pkg.ReplyService.NewestReplies()
	if err != nil {
		return nil, err
	}
	return web.ViewData{"Categories": c, "NewestArticles": a, "NewestReplies": r}, nil
}

// Common 公共列表页面生成方法
func Common(context *web.Context, pkg *ServicePackage, list []*models.ArticleDetail, firstUrl string, baseUrl string, currentPage int, totalPages int) web.GetResult {
	var data, err = layoutViewData(pkg)
	if err != nil {
		context.Processor.Logger.Error(err)
		return context.NotFound()
	}
	data["TotalPages"] = totalPages
	if list != nil {
		data["Articles"] = list
	}
	data["CurrentPage"] = currentPage
	data["FirstUrl"] = firstUrl
	data["BaseUrl"] = baseUrl

	// 计算起止页码 取当前页码的左右3页
	var start = currentPage - 3
	if start <= 0 {
		start = 1
	}
	var end = start + 6
	if end > totalPages {
		end = totalPages
		start = end - 6
		if start <= 0 {
			start = 1
		}
	}
	data["StartPage"] = start
	data["EndPage"] = end
	return context.View("pages/index.html", data)
}

// Index 首页
func Index(context *web.Context, pkg *ServicePackage, param struct {
	Page int `?;>0`
}) web.GetResult {
	var count, err = pkg.ArticleService.AvailableNum()
	if err != nil {
		context.Processor.Logger.Error(err)
		return context.NotFound()
	}
	var totalPages = calculatePages(count)
	var list []*models.ArticleDetail
	if totalPages > 0 {
		if param.Page < 1 {
			param.Page = 1
		}
		list, err = pkg.ArticleService.ListAvailable(param.Page, pageCount)
		if err != nil || len(list) <= 0 {
			context.Processor.Logger.Error(err)
			return context.NotFound()
		}
	}
	return Common(context, pkg, list, "/", "/p", param.Page, totalPages)
}

// Category 分类列表
func Category(context *web.Context, pkg *ServicePackage, param struct {
	Category int `!;>0`
	Page     int `?;>0`
}) web.GetResult {
	var exist, err = pkg.CategoryService.Exist(param.Category)
	if !exist || err != nil {
		context.Processor.Logger.Error(err)
		return context.NotFound()
	}
	count, err := pkg.ArticleService.SpecAvailableNum(param.Category)
	if err != nil {
		context.Processor.Logger.Error(err)
		return context.NotFound()
	}
	var totalPages = calculatePages(count)
	var list []*models.ArticleDetail
	if totalPages > 0 {
		if param.Page < 1 {
			param.Page = 1
		}
		list, err = pkg.ArticleService.ListSpecAvailable(param.Category, param.Page, pageCount)
		if err != nil || len(list) <= 0 {
			context.Processor.Logger.Error(err)
			return context.NotFound()
		}
	}
	var base = "/c" + strconv.Itoa(param.Category)
	return Common(context, pkg, list, base, base+"/p", param.Page, totalPages)
}

// Article 文章
func Article(context *web.Context, pkg *ServicePackage, param struct {
	Article int `!;>0`
}) web.GetResult {
	var data, err = layoutViewData(pkg)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	article, err := pkg.ArticleService.AvailableArticle(param.Article)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	data["Article"] = article
	pre, err := pkg.ArticleService.PreviousArticle(article.Id, article.Category)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	data["PreArticle"] = pre
	next, err := pkg.ArticleService.NextArticle(article.Id, article.Category)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	data["NextArticle"] = next

	count, err := pkg.ReplyService.ReplyNum(article.Id)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	var pages = calculatePages(count)
	data["ReplyPages"] = pages
	if pages > 0 {
		replies, err := pkg.ReplyService.ListAll(article.Id, 1, pageCount)
		if err != nil {
			context.Processor.Logger.Error(err.Error())
			return context.NotFound()
		}
		data["Replies"] = replies
	}
	return context.View("pages/article.html", data)
}

// Manager 返回管理页面
func Manager(context *web.Context) web.GetResult {
	return context.View("manager/index.html")
}
