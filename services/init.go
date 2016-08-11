package services

import (
	"reflect"

	"github.com/kdada/tinygo/meta"
)

func init() {
	//注册服务类型,用于自动注入
	meta.RegisterGenerateFunc(reflect.TypeOf((*UserService)(nil)), func() interface{} {
		return NewUserService()
	})
	meta.RegisterGenerateFunc(reflect.TypeOf((*CategoryService)(nil)), func() interface{} {
		return NewCategoryService()
	})
	meta.RegisterGenerateFunc(reflect.TypeOf((*ArticleService)(nil)), func() interface{} {
		return NewArticleService()
	})
	meta.RegisterGenerateFunc(reflect.TypeOf((*ReplyService)(nil)), func() interface{} {
		return NewReplyService()
	})
}
