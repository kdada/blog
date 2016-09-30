package services

import (
	"math/rand"
	"time"

	"github.com/kdada/tinygo/sql"

	"github.com/kdada/tinygo/meta"
)

func init() {
	//随机数初始化
	rand.Seed(time.Now().UnixNano())
	// 注册服务
	meta.GlobalValueContainer.Register(nil, func() *sql.DB {
		return sql.OpenDefault()
	})
	meta.GlobalValueContainer.Register(nil, (*UserService)(nil))
	meta.GlobalValueContainer.Register(nil, (*CategoryService)(nil))
	meta.GlobalValueContainer.Register(nil, (*FileService)(nil))
	meta.GlobalValueContainer.Register(nil, (*ArticleService)(nil))
	meta.GlobalValueContainer.Register(nil, (*ReplyService)(nil))
}
