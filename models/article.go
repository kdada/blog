package models

import "time"

// 分类
type Category struct {
	Id   int    //分类id
	Name string //分类名称
}

// 文章
type Article struct {
	Id         int       //文章id
	Title      string    //文章标题
	Summary    string    //文章概要
	Text       string    //文章内容
	CreateTime time.Time //创建时间
	UpdateTime time.Time //更新时间
}

func NewArticle(id int, title string) *Article {
	return &Article{
		id,
		title,
		"",
		"",
		time.Now(),
		time.Now(),
	}
}

type Reply struct {
	Article    int       //文章id
	Title      string    //文章名称
	Id         int       //回复id
	Order      int       //回复序号
	Author     int       //作者id
	Name       string    //作者昵称
	Content    string    //回复内容
	CreateTime time.Time //回复时间
}

func NewReply(id int, content string) *Reply {
	return &Reply{
		Id:      id,
		Content: content,
	}
}
