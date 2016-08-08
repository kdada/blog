package service

import "blog/model"

// 文章服务
type ArticleService struct {
}

// Categories 获取所有分类
func (this *ArticleService) Categories() ([]*model.Category, error) {
	return []*model.Category{&model.Category{
		111,
		"魔性分类",
	}}, nil
}

// NewestArticles 获取最新文章
func (this *ArticleService) NewestArticles() ([]*model.Article, error) {
	return []*model.Article{
		model.NewArticle(1, "魔性文章1"),
		model.NewArticle(2, "魔性文章2"),
		model.NewArticle(3, "魔性文章3"),
		model.NewArticle(3, "魔性文章4"),
	}, nil
}

// NewestArticles 获取最新文章
func (this *ArticleService) NewestReplies() ([]*model.Reply, error) {
	return []*model.Reply{
		model.NewReply(1, "魔性回复1"),
		model.NewReply(2, "魔性回复2"),
		model.NewReply(3, "魔性回复3"),
	}, nil
}
