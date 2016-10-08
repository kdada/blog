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
func Common(context *web.Context, pkg *ServicePackage, title string, list []*models.ArticleDetail, firstUrl string, baseUrl string, currentPage int, totalPages int) web.GetResult {
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
	data["Title"] = title
	return context.View("pages/index.html", data)
}

// Index 首页
func Index(context *web.Context, pkg *ServicePackage, params struct {
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
		if params.Page < 1 {
			params.Page = 1
		}
		list, err = pkg.ArticleService.ListAvailable(params.Page, pageCount)
		if err != nil || len(list) <= 0 {
			context.Processor.Logger.Error(err)
			return context.NotFound()
		}
	}
	return Common(context, pkg, "", list, "/", "/p", params.Page, totalPages)
}

// Category 分类列表
func Category(context *web.Context, pkg *ServicePackage, params struct {
	Category int `!;>0`
	Page     int `?;>0`
}) web.GetResult {
	var category, err = pkg.CategoryService.AvailableCategory(params.Category)
	if category == nil || err != nil {
		context.Processor.Logger.Error(err)
		return context.NotFound()
	}
	count, err := pkg.ArticleService.SpecAvailableNum(params.Category)
	if err != nil {
		context.Processor.Logger.Error(err)
		return context.NotFound()
	}
	var totalPages = calculatePages(count)
	var list []*models.ArticleDetail
	if totalPages > 0 {
		if params.Page < 1 {
			params.Page = 1
		}
		list, err = pkg.ArticleService.ListSpecAvailable(params.Category, params.Page, pageCount)
		if err != nil || len(list) <= 0 {
			context.Processor.Logger.Error(err)
			return context.NotFound()
		}
	}
	var base = "/c" + strconv.Itoa(params.Category)
	return Common(context, pkg, category.Name, list, base, base+"/p", params.Page, totalPages)
}

// CalculateReply 计算回复相关内容
func CalculateReply(replyService *services.ReplyService, data web.ViewData, article int, page int) error {
	var count, err = replyService.ReplyNum(article)
	if err != nil {
		return err
	}
	var totalPages = calculatePages(count)
	if page > totalPages {
		page = totalPages
	}
	data["TotalPages"] = totalPages
	if page > 0 {
		replies, err := replyService.ListAll(article, page, pageCount)
		if err != nil {
			return err
		}
		data["Replies"] = replies
	}
	data["CurrentPage"] = page

	// 计算起止页码 取当前页码的左右3页
	var start = page - 3
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
	return nil
}

// Article 文章
func Article(context *web.Context, pkg *ServicePackage, params struct {
	Article int `!;>0`
}) web.GetResult {
	var data, err = layoutViewData(pkg)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	article, err := pkg.ArticleService.AvailableArticle(params.Article)
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
	err = CalculateReply(pkg.ReplyService, data, params.Article, 1)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	return context.View("pages/article.html", data)
}

// ReplyView 显示回复列表指定页
func ReplyView(context *web.Context, replyService *services.ReplyService, params struct {
	Article int `!;>0` //文章id
	Page    int `!;>0` //页码
}) web.PostResult {
	var data = web.ViewData{}
	var err = CalculateReply(replyService, data, params.Article, params.Page)
	if err != nil {
		context.Processor.Logger.Error(err.Error())
		return context.NotFound()
	}
	return context.PartialView("pages/reply.html", data)
}

// Manager 返回管理页面
func Manager(context *web.Context, userService *services.UserService) web.GetResult {
	var obj, ok = context.Session.Value("User")
	if ok {
		var user = obj.(*models.UserDetail)
		var admin, err = userService.IsAdmin(user.Id)
		if err == nil && admin {
			return context.View("manager/index.html")
		}
	}
	return context.NotFound()
}
