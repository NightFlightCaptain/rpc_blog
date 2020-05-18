package service

import (
	"rpc_blog_article/models"
	"rpc_blog_article/pkg/e"
)

type ArticleService struct {
}

type Article struct {
	Id    int
	TagId int

	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int

	CreatedBy  string
	ModifiedBy string
}


func (articleService *ArticleService) ExistArticleById(articleId int) (exists bool, code int) {
	exists, err := models.ExistArticleByID(articleId)
	if err != nil {
		return false, e.ERROR
	}
	if !exists {
		return exists, e.ERROR_NOT_EXIST_ARTICLE
	}
	return exists, e.SUCCESS
}

func (articleService *ArticleService) GetArticle(articleId int) (article models.Article, code int) {
	_, code = articleService.ExistArticleById(articleId)
	if code != e.SUCCESS {
		return
	}
	article, err := models.GetArticle(articleId)
	if err != nil {
		return article, e.ERROR
	}
	return article, e.SUCCESS
}

func (articleService *ArticleService) GetArticles(offset, limit int, maps map[string]interface{}) (map[string]interface{}, int) {
	code := e.ERROR
	maps["deleted_on"] = 0
	articles, err := models.GetArticles(offset, limit, maps)
	if err != nil {
		return nil, code
	}

	count, err := models.GetArticleTotal(maps)
	if err != nil {
		return nil, code
	}

	data := make(map[string]interface{})
	data["list"] = articles
	data["total"] = count

	code = e.SUCCESS
	return data, code
}
//
//func (articleService *ArticleService) AddArticle(article Article) (code int) {
//	data := map[string]interface{}{
//		"tag_id":          article.TagId,
//		"title":           article.Title,
//		"desc":            article.Desc,
//		"content":         article.Content,
//		"created_by":      article.CreatedBy,
//		"cover_image_url": article.CoverImageUrl,
//		"state":           article.State,
//	}
//	_, err := models.AddArticle(data)
//	if err != nil {
//		return e.ERROR
//	}
//	code = e.SUCCESS
//	return
//}
//
//func (articleService *ArticleService) DeleteArticle(articleId int) (code int) {
//	_, code = articleService.ExistArticleById(articleId)
//	if code != e.SUCCESS {
//		return code
//	}
//	_, err := models.DeleteArticle(articleId)
//	if err != nil {
//		return e.ERROR
//	}
//	code = e.SUCCESS
//	return
//}
//
//func (articleService *ArticleService) EditArticle(article Article) (code int) {
//	_, code = articleService.ExistArticleById(article.Id)
//	if code != e.SUCCESS {
//		return code
//	}
//
//	if article.TagId != 0 {
//		_, code = tagService.ExistTagById(article.TagId)
//		if code != e.SUCCESS {
//			return code
//		}
//	}
//
//	data := make(map[string]interface{})
//	data["modified_by"] = article.ModifiedBy
//	if article.TagId != 0 {
//		data["tag_id"] = article.TagId
//	}
//	if article.Title != "" {
//		data["title"] = article.Title
//	}
//	if article.Desc != "" {
//		data["desc"] = article.Desc
//	}
//	if article.Content != "" {
//		data["content"] = article.Content
//	}
//	if article.CoverImageUrl != "" {
//		data["cover_image_url"] = article.CoverImageUrl
//	}
//	if article.State != 0 {
//		data["state"] = article.State
//	}
//	_, err := models.EditArticle(article.Id, data)
//	if err != nil {
//		return e.ERROR
//	}
//	code = e.SUCCESS
//	return
//}
