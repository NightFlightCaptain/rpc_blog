package rpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"rpc_blog_client/models"
	out "rpc_blog_client/rpc/proto"
)

var Client out.TagServiceClient

func SetUp() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	Client = out.NewTagServiceClient(conn)
}

func GetTag(id int) (*models.Tag, int) {
	resp, err := Client.GetTag(context.Background(), &out.Id{Id: int32(id)})
	if err != nil {
		log.Fatalf("couldn't create customer: %v ", err)
	}
	tag := &models.Tag{
		Model:      models.Model{ID: int(resp.Tag.Id)},
		Name:       resp.Tag.Name,
		CreatedBy:  resp.Tag.CreatedBy,
		ModifiedBy: resp.Tag.ModifiedBy,
		State:      int(resp.Tag.State),
	}
	return tag, int(resp.Code)

}
