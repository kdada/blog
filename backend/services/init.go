package services

import (
	"github.com/kdada/tinygo/sql"

	"github.com/kdada/tinygo/meta"
)

// 注册服务
func init() {
	meta.GlobalValueContainer.Register(nil, func() *sql.DB {
		return sql.OpenDefault()
	})
	meta.GlobalValueContainer.Register(nil, (*UserService)(nil))
	meta.GlobalValueContainer.Register(nil, (*CategoryService)(nil))
	meta.GlobalValueContainer.Register(nil, (*FileService)(nil))
}
