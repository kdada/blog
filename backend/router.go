package backend

import (
	"blog/backend/controllers"

	"github.com/kdada/tinygo/router"
	"github.com/kdada/tinygo/web"
)

// routers 创建应用路由
func routers() router.Router {
	//首页路由
	var root = web.NewFuncRouter("", controllers.Index)
	root.AddChild(web.NewFuncRouter(`p{Page=\d+?}.html`, controllers.Index))

	//分类列表路由
	var category = web.NewFuncRouter(`c{Category=\d+}`, controllers.Category)
	category.AddChild(web.NewFuncRouter(`p{Page=\d+}.html`, controllers.Category))
	root.AddChild(category)

	//文章路由
	root.AddChild(web.NewFuncRouter(`a{Article=\d+?}.html`, controllers.Article))

	//非管理员接口
	var reply = web.NewControllerRouter(new(controllers.ReplyController))
	reply.AddChild(web.NewFuncRouter("view", controllers.ReplyView))
	root.AddChild(reply)
	root.AddChild(web.NewControllerRouter(new(controllers.AccountController)))

	//需登陆接口
	var loginFilter = new(LoginFilter)
	var logout, ok = root.Find(router.NewBaseContext("/account/logout"))
	if !ok {
		panic("/account/logout路由不存在")
	}
	logout.AddPreFilter(loginFilter)
	create, ok := root.Find(router.NewBaseContext("/reply/create"))
	if !ok {
		panic("/reply/create路由不存在")
	}
	create.AddPreFilter(loginFilter)
	//管理员可访问部分
	var adminFilter = new(AdminFilter)
	var manager = web.NewSpaceRouter("manager").AddPreFilter(adminFilter)
	manager.AddChild(web.NewMutableFuncRouter("", controllers.Manager))
	root.AddChild(manager)
	root.AddChild(web.NewControllerRouter(new(controllers.UserController)).AddPreFilter(adminFilter))
	root.AddChild(web.NewControllerRouter(new(controllers.CategoryController)).AddPreFilter(adminFilter))
	root.AddChild(web.NewControllerRouter(new(controllers.FileController)).AddPreFilter(adminFilter))
	root.AddChild(web.NewControllerRouter(new(controllers.ArticleController)).AddPreFilter(adminFilter))
	return root
}
