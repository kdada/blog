package models

import (
	"time"
)

// 文章概要
type ArticleSummary struct {
	Id         int       //文章id
	Title      string    //文章标题
	Summary    string    //文章概要
	IsTop      bool      //是否置顶
	CreateTime time.Time //创建时间
	UpdateTime time.Time //更新时间
}

// 文章详细信息
type Article struct {
	Id         int       //文章id
	Category   int       //分类id
	Name       string    //分类名称
	Title      string    //文章标题
	Content    string    //文章内容
	IsTop      bool      //是否置顶
	CreateTime time.Time //创建时间
	UpdateTime time.Time //更新时间
}
