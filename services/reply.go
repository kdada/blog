package services

import (
	"blog/models"

	"github.com/kdada/tinygo/sql"
)

// 评论服务
type ReplyService struct {
	db *sql.DB
}

// NewReplyService 创建评论服务
func NewReplyService() *ReplyService {
	return &ReplyService{
		sql.OpenDefault(),
	}
}

// NewestReplies 获取最新评论
func (this *ReplyService) NewestReplies() ([]*models.ReplyDetail, error) {
	var v []*models.ReplyDetail
	var _, err = this.db.Query(`
select r.id,r.article,r.account,left(r.content,20) as  content,r.create_time,a.title,c.name from 
article a,account c,
(select * 
from reply 
where status = 1 
order by id desc 
limit 10) r
where a.id = r.article and c.id = r.account
`).Scan(&v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

// Replies 获取指定文章指定页的评论,page从1开始
func (this *ReplyService) Replies(article int, page int) ([]*models.Reply, error) {
	var v []*models.Reply
	var _, err = this.db.Query(`
select r.*,u.name from account u,
(select *,row_number() over(order by id asc) as floor
from reply
where article = $1
order by id asc) r
where r.status = 1 and u.id = r.account
order by floor desc
offset $2
limit $3
`, article, (page-1)*10, 10).Scan(&v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

// Page 获取指定文章评论的总页数
func (this *ReplyService) Page(article int) (int, error) {
	var v int
	var _, err = this.db.Query(`select count(*) as floor from reply where article = $1 and status = 1`, article).Scan(&v)
	if err == nil {
		var page = v / 10
		if v%10 > 0 {
			page++
		}
		return page, nil
	}
	return 0, err
}

// New 创建评论
func (this *ReplyService) New(article, reply, userId int, content string) error {
	var _, err = this.db.Exec(`insert into reply(article,account,reply,content) values($1,$2,$3,$4)`, article, userId, reply, content)
	return err
}
