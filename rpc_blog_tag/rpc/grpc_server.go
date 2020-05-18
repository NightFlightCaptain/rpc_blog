package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"rpc_blog_tag/models"
	out "rpc_blog_tag/rpc/proto"
	"rpc_blog_tag/service"
)

type TagServer struct {
}

var tagService = &service.TagService{}

func (s *TagServer) ExistTagById(ctx context.Context, id *out.Id) (*out.ExistResponse, error) {
	exists, code := tagService.ExistTagById(int(id.Id))
	return &out.ExistResponse{
		Exist: exists,
		Code:  int32(code),
	}, nil
}

func (s *TagServer) ExistTagByName(ctx context.Context, name *out.Name) (*out.ExistResponse, error) {
	exists, code := tagService.ExistTagByName(name.Name)
	return &out.ExistResponse{
		Exist: exists,
		Code:  int32(code),
	}, nil
}

func (s *TagServer) GetTag(ctx context.Context, id *out.Id) (*out.GetTagResponse, error) {
	tag, code := tagService.GetTag(int(id.Id))
	return &out.GetTagResponse{
		Tag: &out.Tag{
			Id:         int32(tag.ID),
			Name:       tag.Name,
			CreatedBy:  tag.CreatedBy,
			ModifiedBy: tag.ModifiedBy,
			State:      int32(tag.State),
		},
		Code: int32(code),
	}, nil
}

func getOutTagFromModelTag(in models.Tag) (tag *out.Tag) {
	tag = &out.Tag{}
	tag.Id = int32(in.ID)
	tag.Name = in.Name
	tag.State = int32(in.State)
	tag.ModifiedBy = in.ModifiedBy
	tag.CreatedBy = in.CreatedBy
	return
}

func getModelsTagFromOutTag(in *out.Tag) (tag models.Tag) {
	tag.Model.ID = int(in.Id)
	tag.Name = in.Name
	tag.State = int(in.State)
	tag.ModifiedBy = in.ModifiedBy
	tag.CreatedBy = in.CreatedBy
	return
}

func (s *TagServer) GetTags(ctx context.Context, in *out.GetTagsRequest) (*out.GetTagsResponse, error) {
	maps := make(map[string]interface{})
	fmt.Println(in.Maps)
	err := json.Unmarshal(in.Maps, &maps)
	if err != nil {
		log.Println("json unmarshal error:", err)
		return nil, errors.New("json error")
	}
	data, code := tagService.GetTags(int(in.Offset), int(in.Limit), maps)

	modelsTags := data["list"].([]models.Tag)
	tags := make([]*out.Tag, 0,len(modelsTags))
	for _, t := range modelsTags {
		tags = append(tags, getOutTagFromModelTag(t))
	}

	resData := &out.GetTagsResponse_Data{
		Tags:  tags,
		Total: int32(data["total"].(int)),
	}
	return &out.GetTagsResponse{
		Data: resData,
		Code: int32(code),
	}, nil
}

func (s *TagServer) AddTag(ctx context.Context, tag *out.Tag) (*out.Code, error) {
	modelsTag := getModelsTagFromOutTag(tag)
	code := tagService.AddTag(modelsTag)
	return &out.Code{Code: int32(code)}, nil
}

func (s *TagServer) EditTag(ctx context.Context, tag *out.Tag) (*out.Code, error) {
	modelsTag := getModelsTagFromOutTag(tag)
	code := tagService.EditTag(modelsTag)
	return &out.Code{Code: int32(code)}, nil
}

func (s *TagServer) DeleteTag(ctx context.Context, tag *out.Id) (*out.Code, error) {
	code := tagService.DeleteTag(int(tag.Id))
	return &out.Code{Code: int32(code)}, nil
}
