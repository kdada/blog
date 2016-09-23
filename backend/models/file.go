package models

import "time"

// 文件详细信息
type FileDetail struct {
	Id         int       //文件id
	FileName   string    //文件名
	Content    string    //文件描述
	UploadTime time.Time //上传时间
	Status     int       //状态码:1-正常,2-已删除
}
