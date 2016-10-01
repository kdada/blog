package services

import (
	"blog/backend/models"

	"github.com/kdada/tinygo/sql"
)

// 评论服务
type ReplyService struct {
	DB *sql.DB
}

// NewReplyService 创建评论服务
func NewReplyService() *ReplyService {
	return &ReplyService{
		sql.OpenDefault(),
	}
}

// NewestReplies 获取最新评论
func (this *ReplyService) NewestReplies() (replies []*models.ReplyDetail, err error) {
	_, err = this.DB.Query(`
select r.id,r.article,r.account,substr(regexp_replace(r.content,'<br/>.*$',''),0,50) as  content,r.create_time,a.title,c.name from
article a,account c,
(select *
from reply
where status = 1
order by id desc
limit 10) r
where a.id = r.article and c.id = r.account
order by r.id desc
`).Scan(&replies)
	return
}

// ListAll 获取指定文章指定页的评论
func (this *ReplyService) ListAll(article int, page int, count int) (replies []*models.Reply, err error) {
	_, err = this.DB.Query(`
select r.*,u.name from account u,
(select *,row_number() over(order by id asc) as floor
from reply
where article = $1
order by id asc) r
where r.status = 1 and u.id = r.account
order by floor desc
limit $2
offset $3
`, article, count, (page-1)*count).Scan(&replies)
	return
}

// Page 获取指定文章评论的总页数
func (this *ReplyService) ReplyNum(article int) (count int, err error) {
	_, err = this.DB.Query(`select count(*) as floor from reply where article = $1 and status = 1`, article).Scan(&count)
	return
}

// Create 创建评论
func (this *ReplyService) Create(article, reply, userId int, content string) error {
	return this.DB.Exec(`insert into reply(article,account,reply,content) values($1,$2,$3,$4)`, article, userId, reply, content).Error()
}
