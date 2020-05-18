package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"rpc_blog_article/models"
	"rpc_blog_article/pkg/e"
	"rpc_blog_article/rpc/out"
	"rpc_blog_article/service"
)

type ArticleServer struct {
}

var articleService = &service.ArticleService{}

func getOutArticleFromModelArticle(a models.Article) *out.Article {
	return &out.Article{
		Id:            int32(a.ID),
		Title:         a.Title,
		Desc:          a.Desc,
		Content:       a.Content,
		CoverImageUrl: a.CoverImageUrl,
		CreatedBy:     a.CreatedBy,
		ModifiedBy:    a.ModifiedBy,
		State:         int32(a.State),
		TagId:         int32(a.TagID),
		Tag:           nil,
	}
}

func getModelArticleFromOutArticle(a *out.Article) *models.Article {
	return &models.Article{
		Model: models.Model{
			ID: int(a.Id),
		},
		TagID:         int(a.TagId),
		Title:         a.Title,
		Desc:          a.Desc,
		Content:       a.Content,
		CoverImageUrl: a.CoverImageUrl,
		CreatedBy:     a.CreatedBy,
		ModifiedBy:    a.ModifiedBy,
		State:         int(a.State),
	}
}

func (a ArticleServer) GetArticle(ctx context.Context, in *out.Id) (*out.GetArticleResponse, error) {
	article, code := articleService.GetArticle(int(in.Id))
	tag, err := getTag(ctx, article.TagID)
	if err != nil {
		log.Println("getTag err,", err)
		return nil, err
	}
	outArticle := getOutArticleFromModelArticle(article)
	outArticle.Tag = tag
	return &out.GetArticleResponse{
		Article: outArticle,
		Code:    int32(code),
	}, nil
}

func getTag(ctx context.Context, id int) (*out.Tag, error) {
	resp, err := TagClient.GetTag(ctx, &out.Id{Id: int32(id)})
	if err != nil {
		log.Println("tagclient get tag,",err)
		return nil, err
	}
	if resp.Code != e.SUCCESS {
		return nil, errors.New("get tag error")
	}
	return resp.Tag, nil
}

func (a ArticleServer) GetArticles(ctx context.Context, in *out.GetArticlesRequest) (*out.GetArticlesResponse, error) {
	maps := make(map[string]interface{})
	err := json.Unmarshal(in.Maps, maps)
	if err != nil {
		log.Println("json err,", err)
		return nil, err
	}
	data, code := articleService.GetArticles(int(in.Offset), int(in.Limit), maps)

	modelsArticles := data["list"].([]models.Article)
	articles := make([]*out.Article, 0, len(modelsArticles))
	for _, t := range modelsArticles {
		articles = append(articles, getOutArticleFromModelArticle(t))
	}
	resData := &out.GetArticlesResponse_Data{
		Articles: articles,
		Total:    int32(data["total"].(int)),
	}

	return &out.GetArticlesResponse{
		Data: resData,
		Code: int32(code),
	}, nil
}

func (a ArticleServer) AddArticle(ctx context.Context, in *out.Article) (*out.Code, error) {
	panic("implement me")
}

func (a ArticleServer) EditArticle(ctx context.Context, in *out.Article) (*out.Code, error) {
	panic("implement me")
}

func (a ArticleServer) DeleteArticle(ctx context.Context, in *out.Id) (*out.Code, error) {
	panic("implement me")
}
