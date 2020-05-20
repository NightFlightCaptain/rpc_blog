package rpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"rpc_blog_client/pkg/e"
	"rpc_blog_client/rpc/out"
)

var ArticleClient out.ArticleServiceClient

func setUpArticle(target string) {

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	ArticleClient = out.NewArticleServiceClient(conn)
}

func GetArticle(id int) (*out.Article, int) {
	resp, err := ArticleClient.GetArticle(context.Background(), &out.Id{Id: int32(id)})
	if err != nil {
		log.Println("gerArticle", err)
		return nil, e.ERROR
	}
	return resp.Article, int(resp.Code)
}
