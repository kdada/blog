package services

import "blog/models"

// 文章服务
type ArticleService struct {
}

// Categories 获取所有分类
func (this *ArticleService) Categories() ([]*models.Category, error) {
	return []*models.Category{&models.Category{
		111,
		"魔性分类",
	}}, nil
}

// NewestArticles 获取最新文章
func (this *ArticleService) NewestArticles() ([]*models.Article, error) {
	return []*models.Article{
		models.NewArticle(1, "魔性文章1"),
		models.NewArticle(2, "魔性文章2"),
		models.NewArticle(3, "魔性文章3"),
		models.NewArticle(3, "魔性文章4"),
	}, nil
}

// NewestArticles 获取最新文章
func (this *ArticleService) NewestReplies() ([]*models.Reply, error) {
	return []*models.Reply{
		models.NewReply(1, "魔性回复1"),
		models.NewReply(2, "魔性回复2"),
		models.NewReply(3, "魔性回复3"),
	}, nil
}
