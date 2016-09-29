package models

import (
	"time"
)

// 文章创建信息
type ArticleInfo struct {
	Category int    `!;>0`                //文章分类Id
	Title    string `!;clen>=2&&clen<=20` //文章标题
	Content  string `!;clen>=2`           //文章内容
}

// 文章修改信息
type ArticleData struct {
	ArticleInfo
	Id int `!;>0` //文章id
}

// 文章概要
type ArticleDetail struct {
	Id         int       //文章id
	Title      string    //文章标题
	Content    string    //文章内容
	Category   int       //文章分类Id
	Name       string    //文章分类名称
	Top        int       //文章置顶时间戳
	CreateTime time.Time //文章创建时间
	UpdateTime time.Time //文章最后更新时间
	Status     int       //文章状态
}

// IsTop 返回是否置顶
func (this *ArticleDetail) IsTop() bool {
	return this.Top > 0
}
