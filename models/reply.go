package models

import "time"

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
