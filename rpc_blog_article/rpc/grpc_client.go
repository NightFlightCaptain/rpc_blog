package rpc

import (
	"google.golang.org/grpc"
	"log"
	"rpc_blog_article/rpc/out"
)

var TagClient out.TagServiceClient

func SetUp() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	TagClient = out.NewTagServiceClient(conn)
}
