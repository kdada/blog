package controllers

import (
	"blog/backend/services"

	"github.com/kdada/tinygo/web"
)

// 文件管理控制器
type FileController struct {
	BaseController
	FileService *services.FileService
}

// List 显示列表,每页10行
func (this *FileController) List(params struct {
	Page int `!;>0` //页码
}) web.PostResult {
	return this.returnPostResult(this.FileService.ListAll(params.Page, pageCount))
}

// ListNum 返回页数
func (this *FileController) ListNum() web.PostResult {
	var count, err = this.FileService.FileNum()
	var result interface{} = nil
	if err == nil {
		result = map[string]int{
			"Count": calculatePages(count),
		}
	}
	return this.returnPostResult(result, err)
}

// Upload 上传文件
func (this *FileController) Upload(params struct {
	File *web.FormFile `!`
}) web.PostResult {
	var name, err = this.FileService.Save(params.File)
	var result interface{} = nil
	if err == nil {
		result = map[string]string{
			"Name": name,
		}
	}
	return this.returnPostResult(result, err)
}

// Delete 删除文件
func (this *FileController) Delete(params struct {
	File int `!;>0` //id
}) web.PostResult {
	return this.returnPostResult(nil, this.FileService.Delete(params.File))
}
