package models

import "time"

// 分类详细信息
type CategoryDetail struct {
	Id         int       //分类id
	Name       string    //分类名称
	CreateTime time.Time //创建时间
	Status     int       //状态码:1-正常,2-隐藏,3-删除
}
