package controllers

import (
	"blog/backend/services"

	"github.com/kdada/tinygo/web"
)

// 分类管理控制器
type CategoryController struct {
	BaseController
	CategoryService *services.CategoryService
}

// List 显示分类列表,每页10行
func (this *CategoryController) List(params struct {
	Page int `!;>0` //页码
}) web.PostResult {
	return this.returnPostResult(this.CategoryService.ListAll(params.Page, pageCount))
}

// ListNum 返回页数
func (this *CategoryController) ListNum() web.PostResult {
	var count, err = this.CategoryService.CategoryNum()
	var result interface{} = nil
	if err == nil {
		result = map[string]int{
			"Count": calculatePages(count),
		}
	}
	return this.returnPostResult(result, err)
}

// Create 创建分类
func (this *CategoryController) Create(params struct {
	Name string `!;len>=2&&len<=50` //分类名称
}) web.PostResult {
	return this.returnPostResult(nil, this.CategoryService.Create(params.Name))
}

// Delete 删除分类
func (this *CategoryController) Delete(params struct {
	Category int `!;>0` //分类id
}) web.PostResult {
	return this.returnPostResult(nil, this.CategoryService.Delete(params.Category))
}

// Hide 隐藏分类
func (this *CategoryController) Hide(params struct {
	Category int `!;>0` //分类id
}) web.PostResult {
	return this.returnPostResult(nil, this.CategoryService.Hide(params.Category))
}

// Show 显示分类
func (this *CategoryController) Show(params struct {
	Category int `!;>0` //分类id
}) web.PostResult {
	return this.returnPostResult(nil, this.CategoryService.Show(params.Category))
}
