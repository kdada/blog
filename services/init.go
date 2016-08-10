package services

import "github.com/kdada/tinygo/meta"

func init() {
	//注册服务类型,用于自动注入
	meta.RegisterInstance((*UserService)(nil))
	meta.RegisterInstance((*ArticleService)(nil))
}
