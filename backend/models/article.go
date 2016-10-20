package models

import (
	"time"
	"unicode/utf8"
)

// 文章创建信息
type ArticleInfo struct {
	Category int    `!;>0`                 //文章分类Id
	Title    string `!;clen>=2&&clen<=200` //文章标题
	Content  string `!;clen>=2`            //文章内容
	Summary  string //文章概览
	Html     string //文章html内容
}

// Validate 验证字段
func (this *ArticleInfo) Validate() bool {
	var titleLen = utf8.RuneCountInString(this.Title)
	return this.Category > 0 && titleLen >= 2 && titleLen <= 200 && utf8.RuneCountInString(this.Content) >= 2
}

// 文章修改信息
type ArticleData struct {
	ArticleInfo
	Id int `!;>0` //文章id
}

// Validate 验证字段
func (this *ArticleData) Validate() bool {
	return this.Id > 0 && this.ArticleInfo.Validate()
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
