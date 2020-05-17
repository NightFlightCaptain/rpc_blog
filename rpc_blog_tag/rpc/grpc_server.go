package rpc

import (
	"context"
	out "rpc_blog/rpc/proto"
	"rpc_blog/service"
)

type TagServer struct {
}

var tagService = &service.TagService{}

func (s *TagServer) ExistTagById(ctx context.Context, id *out.Id) (*out.ExistResponse, error) {
	exist, code := tagService.ExistTagById(int(id.Id))
	return &out.ExistResponse{
		Exist: exist,
		Code:  int32(code),
	}, nil
}

func (s *TagServer) ExistTagByName(context.Context, *out.Name) (*out.ExistResponse, error) {
	panic("implement me")
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

func (s *TagServer) GetTags(context.Context, *out.GetTagsRequest) (*out.GetTagsResponse, error) {
	panic("implement me")
}

func (s *TagServer) AddTag(context.Context, *out.Tag) (*out.Code, error) {
	panic("implement me")
}

func (s *TagServer) EditTag(context.Context, *out.Tag) (*out.Code, error) {
	panic("implement me")
}

func (s *TagServer) DeleteTag(context.Context, *out.Id) (*out.Code, error) {
	panic("implement me")
}
