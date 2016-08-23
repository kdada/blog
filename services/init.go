package services

import (
	"github.com/kdada/tinygo/meta"
	"github.com/kdada/tinygo/sql"
)

func init() {
	//注册服务类型,用于自动注入
	meta.GlobalValueContainer.Register(nil, func() *sql.DB {
		return sql.OpenDefault()
	})
	meta.GlobalValueContainer.Register(nil, (*UserService)(nil))
	meta.GlobalValueContainer.Register(nil, (*CategoryService)(nil))
	meta.GlobalValueContainer.Register(nil, (*ArticleService)(nil))
	meta.GlobalValueContainer.Register(nil, (*ReplyService)(nil))
}
