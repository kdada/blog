package services

import (
	"blog/models"

	"github.com/kdada/tinygo/sql"
)

// 回复服务
type ReplyService struct {
	db *sql.DB
}

// NewReplyService 创建回复服务
func NewReplyService() *ReplyService {
	return &ReplyService{
		sql.OpenDefault(),
	}
}

// NewestReplies 获取最新回复
func (this *ReplyService) NewestReplies() ([]*models.Reply, error) {
	var v []*models.Reply
	var _, err = this.db.Query("select * from reply where status = 1 order by id desc limit 10").Scan(&v)
	if err == nil {
		return v, nil
	}
	return nil, err
}
