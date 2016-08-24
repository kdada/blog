package main

import (
	"blog/controllers"
	"blog/models"
	"blog/services"
	"reflect"

	"github.com/kdada/tinygo/meta"
	"github.com/kdada/tinygo/router"
	"github.com/kdada/tinygo/web"
)

// 创建路由
func Router() router.Router {
	//根路由
	var root = web.NewRootRouter()
	//首页路由
	root.AddChild(web.NewFuncRouter(`index.html`, controllers.Index))
	root.AddChild(web.NewFuncRouter(`p{Page=\d+}.html`, controllers.Index))
	//文章路由
	root.AddChild(web.NewFuncRouter(`a{Id=\d+}.html`, controllers.Article))
	//分类列表路由
	var category = web.NewFuncRouter(`c{Id=\d+}`, controllers.Category)
	category.AddChild(web.NewFuncRouter(`p{Page=\d+}.html`, controllers.Category))
	root.AddChild(category)
	//图片路由
	root.AddChild(web.NewStaticRouter("img", "app/img"))
	//管理后台路由
	var manager = web.NewSpaceRouter("manager")
	manager.AddChild(web.NewMutableFuncRouter("index", controllers.Manager))
	manager.AddPreFilter(new(AdminFilter))
	root.AddChild(manager)
	//用户服务路由
	root.AddChild(web.NewControllerRouter(new(controllers.UserController)))
	var logout, _ = root.Find(router.NewBaseContext("/user/logout"))
	logout.AddPreFilter(new(LoginFilter))
	//评论路由
	root.AddChild(web.NewControllerRouter(new(controllers.ReplyController)))
	var newReply, _ = root.Find(router.NewBaseContext("/reply/new"))
	newReply.AddPreFilter(new(LoginFilter))
	//分类路由
	root.AddChild(web.NewControllerRouter(new(controllers.CategoryController)))
	//文章路由
	root.AddChild(web.NewControllerRouter(new(controllers.ArticleController)))
	return root
}

// 登录过滤器
type LoginFilter struct {
}

// Filter 过滤该请求
// return:返回true表示继续处理,否则终止路由过程,后续的过滤器也不会执行
func (this *LoginFilter) Filter(context router.RouterContext) bool {
	var ctx, ok = context.(*web.Context)
	if ok {
		var login, ok = ctx.Session.Bool("Login")
		return ok && login
	}
	return false
}

// 管理员过滤器
type AdminFilter struct {
}

// Filter 过滤该请求
// return:返回true表示继续处理,否则终止路由过程,后续的过滤器也不会执行
func (this *AdminFilter) Filter(context router.RouterContext) bool {
	var ctx, ok = context.(*web.Context)
	if ok {
		var login, ok = ctx.Session.Bool("Login")
		if ok && login {
			var value, _ = ctx.Session.Value("User")
			var loginInfo, ok2 = value.(*models.UserInfo)
			if ok2 {
				var vp, ok3 = meta.GlobalValueContainer.Contains("", reflect.TypeOf((*services.UserService)(nil)))
				if ok3 {
					var userService, ok4 = vp.Value().(*services.UserService)
					return ok4 && userService.IsAdmin(loginInfo.Id)
				}
			}
		}
	}
	return false
}
