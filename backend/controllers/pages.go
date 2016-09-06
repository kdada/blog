package controllers

import "github.com/kdada/tinygo/web"

// Index 返回首页页面
func Index(context *web.Context) web.GetResult {
	return context.View("pages/index.html")
}
