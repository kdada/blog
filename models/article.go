package models

import (
	"time"
)

// 文章
type Article struct {
	Id         int       //文章id
	Title      string    //文章标题
	Summary    string    //文章概要
	Text       string    //文章内容
	CreateTime time.Time //创建时间
	UpdateTime time.Time //更新时间
}
