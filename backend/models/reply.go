package models

import "time"

// 评论记录
type Reply struct {
	Id         int       //评论id
	Account    int       //评论者id
	Name       string    //评论者昵称
	Content    string    //评论内容
	CreateTime time.Time //评论时间
	Floor      int       //楼层
}

// 评论详细记录
type ReplyDetail struct {
	Reply
	Article int    //文章id
	Title   string //文章名称
}
