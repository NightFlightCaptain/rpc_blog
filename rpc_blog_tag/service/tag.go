package service

import (
	"rpc_blog_tag/models"
	"rpc_blog_tag/pkg/e"
)

type TagService struct {
}

func (tagService *TagService) ExistTagById(id int) (exists bool, code int) {
	exists, err := models.ExistTagById(id)
	if err != nil {
		return false, e.ERROR
	}
	if !exists {
		return false, e.ERROR_NOT_EXIST_TAG
	}
	return true, e.SUCCESS
}

/*
ExistTagByName 判断tag名称是否已经存在，如果存在，则返回true和e.ERROR_EXIST_TAG的错误。
如果不存在返回false和e.Success
*/
func (tagService *TagService) ExistTagByName(name string) (exists bool, code int) {
	exists, err := models.ExistTagByName(name)
	if err != nil {
		return false, e.ERROR
	}
	if exists {
		return true, e.ERROR_EXIST_TAG
	}
	return false, e.SUCCESS
}

func (tagService *TagService) GetTag(id int) (tag models.Tag, code int) {
	_, code = tagService.ExistTagById(id)
	if code != e.SUCCESS {
		return
	}

	tag, err := models.GetTag(id)
	if err != nil {
		return tag, e.ERROR
	}
	return tag, e.SUCCESS
}

func (tagService *TagService) GetTags(offset, limit int, maps map[string]interface{}) (map[string]interface{}, int) {
	maps["deleted_on"] = 0
	tags, err := models.GetTags(offset, limit, maps)
	if err != nil {
		return nil, e.ERROR
	}

	//total is the num regardless of limit and offset
	total, err := models.GetTagTotal(maps)
	if err != nil {
		return nil, e.ERROR
	}

	data := map[string]interface{}{
		"list":  tags,
		"total": total,
	}
	return data, e.SUCCESS
}

func (tagService *TagService) AddTag(tag models.Tag) (code int) {
	_, code = tagService.ExistTagByName(tag.Name)
	if code != e.SUCCESS {
		return code
	}

	data := map[string]interface{}{
		"name":       tag.Name,
		"created_by": tag.CreatedBy,
		"state":      1,
	}
	_, err := models.AddTag(data)
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

func (tagService *TagService) EditTag(tag models.Tag) (code int) {
	_, code = tagService.ExistTagById(tag.Model.ID)
	if code != e.SUCCESS {
		return code
	}

	_, code = tagService.ExistTagByName(tag.Name)
	if code != e.SUCCESS {
		return code
	}

	data := make(map[string]interface{})
	data["modified_by"] = tag.ModifiedBy
	if tag.Name != "" {
		data["name"] = tag.Name
	}
	if tag.State != 0 {
		data["state"] = tag.State
	}
	_, err := models.EditTag(tag.Model.ID, data)
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

func (tagService *TagService) DeleteTag(id int) (code int) {
	_, code = tagService.ExistTagById(id)
	if code != e.SUCCESS {
		return code
	}

	_, err := models.DeleteTag(id)
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}
