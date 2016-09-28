package controllers

import (
	"blog/backend/models"

	"github.com/kdada/tinygo/web"
)

// 所有列表每页数量
const pageCount = 10

// calculatePages 根据count数量计算Page页数
func calculatePages(count int) int {
	var pages = count / pageCount
	if count%pageCount > 0 {
		pages++
	}
	return pages
}

// 基础控制器
type BaseController struct {
	Context *web.Context
}

// 返回Post结果
func (this *BaseController) returnPostResult(obj interface{}, err error) web.PostResult {
	if err != nil {
		this.Context.Processor.Logger.Error(err)
		return this.Context.Json(models.NewErrorResult(err))
	}
	return this.Context.Json(models.NewSuccessResult(obj))
}
