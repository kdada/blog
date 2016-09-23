package services

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"blog/backend/models"

	"github.com/kdada/tinygo/sql"
	"github.com/kdada/tinygo/web"
)

const filePath = "./upload/files/"

// 文件服务
type FileService struct {
	DB sql.DB
}

// FileNum 返回文件总数
func (this *FileService) FileNum() (count int, err error) {
	_, err = this.DB.Query("select count(*) from file where status = 1").Scan(&count)
	return
}

// List 返回文件列表
func (this *FileService) List(page int, count int) (details []*models.FileDetail, err error) {
	_, err = this.DB.Query("select * from file where status = 1 limit $1 offset $2", count, (page-1)*count).Scan(&details)
	return
}

// Save 保存文件并返回保存后的文件名称
func (this *FileService) Save(file *web.FormFile) (string, error) {
	var ext = filepath.Ext(file.FileName())
	var fileName = strconv.FormatInt(time.Now().UnixNano(), 10) + ext
	var err = this.DB.Begin()
	if err != nil {
		return "", err
	}
	err = this.DB.Exec("insert into file(file_name,content) values($1,$2)", fileName, file.FileName()).Error()
	if err != nil {
		this.DB.Rollback()
		return "", err
	}
	err = file.SaveTo(filePath + fileName)
	if err != nil {
		this.DB.Rollback()
		return "", err
	}
	err = this.DB.Commit()
	if err != nil {
		return "", err
	}
	return fileName, nil
}

// Delete 删除文件
func (this *FileService) Delete(fileId int) error {
	var f *models.FileDetail
	var count, err = this.DB.Query("select * from file where id = $1 and status = 1", fileId).Scan(&f)
	if err != nil {
		return err
	}
	if count <= 0 {
		return nil
	}
	err = this.DB.Begin()
	if err != nil {
		return err
	}
	err = this.DB.Exec("update file set status = 2 where id = $1", f.Id).Error()
	if err != nil {
		this.DB.Rollback()
		return err
	}
	err = os.Remove(filePath + f.FileName)
	if err != nil {
		this.DB.Rollback()
		return err
	}
	err = this.DB.Commit()
	if err != nil {
		return err
	}
	return nil
}
