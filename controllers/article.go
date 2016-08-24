package controllers

import (
	"blog/models"
	"blog/services"
	"fmt"

	"github.com/kdada/tinygo/web"
)

// 文章控制器
type ArticleController struct {
	Context        *web.Context
	ArticleService *services.ArticleService
}

// 创建文章
func (this *ArticleController) New(param struct {
	Title    string `!,clen>0&&clen<=50`
	Content  string `!,clen>0`
	Category int    `!,>0`
}) web.PostResult {
	fmt.Println(param.Title)
	fmt.Println(param.Content)
	fmt.Println(param.Category)
	var err = this.ArticleService.New(param.Title, param.Content, param.Category)
	if err != nil {
		this.Context.Api(models.NewErrorResult(err))
	}
	return this.Context.Api(models.NewSuccessResult(nil))
}
