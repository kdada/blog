package controllers

import (
	"blog/backend/models"
	"blog/backend/services"

	"github.com/kdada/tinygo/web"
)

// 分类管理控制器
type CategoryManagerController struct {
	Context         *web.Context
	CategoryService *services.CategoryService
}

// List 显示分类列表,每页10行
func (this *CategoryManagerController) List(params struct {
	Page int `!;>0` //页码
}) web.PostResult {
	var r, err = this.CategoryService.ListAll(params.Page, pageCount)
	if err != nil {
		this.Context.Processor.Logger.Error(err)
		return this.Context.Json(models.NewErrorResult(err))
	}
	return this.Context.Json(models.NewSuccessResult(r))
}

// ListNum 返回页数
func (this *CategoryManagerController) ListNum() web.PostResult {
	var r, err = this.CategoryService.CategoryNum()
	if err != nil {
		this.Context.Processor.Logger.Error(err)
		return this.Context.Json(models.NewErrorResult(err))
	}
	var pages = r / pageCount
	if r%pageCount > 0 {
		pages++
	}
	return this.Context.Json(models.NewSuccessResult(map[string]int{
		"Count": pages,
	}))
}

// Create 创建分类
func (this *CategoryManagerController) Create(params struct {
	Name string `!;len>=2&&len<=50` //分类名称
}) web.PostResult {
	var err = this.CategoryService.Create(params.Name)
	if err != nil {
		this.Context.Processor.Logger.Error(err)
		this.Context.Json(models.NewErrorResult(err))
	}
	return this.Context.Json(models.NewSuccessResult(nil))
}

// Delete 删除分类
func (this *CategoryManagerController) Delete(params struct {
	Category int `!;>0` //分类id
}) web.PostResult {
	var err = this.CategoryService.Delete(params.Category)
	if err != nil {
		this.Context.Processor.Logger.Error(err)
		this.Context.Json(models.NewErrorResult(err))
	}
	return this.Context.Json(models.NewSuccessResult(nil))
}

// Hide 隐藏分类
func (this *CategoryManagerController) Hide(params struct {
	Category int `!;>0` //分类id
}) web.PostResult {
	var err = this.CategoryService.Hide(params.Category)
	if err != nil {
		this.Context.Processor.Logger.Error(err)
		this.Context.Json(models.NewErrorResult(err))
	}
	return this.Context.Json(models.NewSuccessResult(nil))
}

// Show 显示分类
func (this *CategoryManagerController) Show(params struct {
	Category int `!;>0` //分类id
}) web.PostResult {
	var err = this.CategoryService.Show(params.Category)
	if err != nil {
		this.Context.Processor.Logger.Error(err)
		this.Context.Json(models.NewErrorResult(err))
	}
	return this.Context.Json(models.NewSuccessResult(nil))
}
