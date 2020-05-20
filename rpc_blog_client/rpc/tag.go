package rpc

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"log"
	"rpc_blog_client/models"
	"rpc_blog_client/pkg/e"
	out "rpc_blog_client/rpc/out"
)

var TagClient out.TagServiceClient

func setUpTag(target string) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	TagClient = out.NewTagServiceClient(conn)
}

func ExistTagById(id int) int {
	resp, err := TagClient.ExistTagById(context.Background(), &out.Id{Id: int32(id)})
	if err != nil {
		log.Printf("error %v", err)
		return e.ERROR
	}
	return int(resp.Code)
}

func ExistTagByName(name string) int {
	resp, err := TagClient.ExistTagByName(context.Background(), &out.Name{Name: name})
	if err != nil {
		log.Printf("error %v", err)
		return e.ERROR
	}
	return int(resp.Code)
}

func GetTag(id int) (*models.Tag, int) {
	resp, err := TagClient.GetTag(context.Background(), &out.Id{Id: int32(id)})
	if err != nil {
		log.Printf("couldn't create customer: %v ", err)
		return nil, e.ERROR
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

func GetTags(offset, limit int, maps map[string]interface{}) (data map[string]interface{}, code int) {
	mapsBytes, err := json.Marshal(maps)
	if err != nil {
		log.Println("json error", err)
	}
	resp, err := TagClient.GetTags(
		context.Background(),
		&out.GetTagsRequest{
			Offset: int32(offset),
			Limit:  int32(limit),
			Maps:   mapsBytes,
		},
	)
	if err != nil {
		log.Println("error,", err)
		return nil, e.ERROR
	}
	data = make(map[string]interface{})
	data["list"] = resp.Data.Tags
	data["total"] = resp.Data.Total
	code = int(resp.Code)
	return
}

func AddTag(tag models.Tag) (code int) {
	resp, err := TagClient.AddTag(context.Background(), &out.Tag{
		Name:      tag.Name,
		CreatedBy: tag.CreatedBy,
		State:     int32(tag.State),
	})
	if err != nil {
		log.Println(err)
	}
	return int(resp.Code)
}

func EditTag(tag models.Tag) (code int) {
	resp, err := TagClient.EditTag(context.Background(), &out.Tag{
		Id:         int32(tag.ID),
		Name:       tag.Name,
		ModifiedBy: tag.ModifiedBy,
		State:      int32(tag.State),
	})
	if err != nil {
		log.Println(err)
	}
	return int(resp.Code)
}

func DeleteTag(id int) (code int) {
	resp, err := TagClient.DeleteTag(context.Background(), &out.Id{Id: int32(id)})
	if err != nil {
		log.Println(err)
	}
	return int(resp.Code)
}
