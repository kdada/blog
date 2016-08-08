package services

import "github.com/kdada/tinygo/meta"

func init() {
	meta.RegisterInstance((*UserService)(nil))
	meta.RegisterInstance((*ArticleService)(nil))
}
