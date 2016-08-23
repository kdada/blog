package controllers

import (
	"blog/models"
	"blog/services"

	"github.com/kdada/tinygo/web"
)

// 分类控制器
type CategoryController struct {
	Context         *web.Context
	CategoryService *services.CategoryService
}

// List 获取分类列表
func (this *CategoryController) List() web.PostResult {
	var c, err = this.CategoryService.Categories()
	if err != nil {
		this.Context.Api(models.NewErrorResult(err))
	}
	return this.Context.Api(models.NewSuccessResult(c))
}
