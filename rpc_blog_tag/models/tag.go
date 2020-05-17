package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTag(id int) (tag Tag, err error) {
	err = db.Where("id = ? and deleted_on = ?", id, 0).First(&tag).Error
	if err != nil {
		log.Printf("models GetTag err:%v", err)
	}
	return
}

func GetTags(offset, limit int, maps interface{}) (tags []Tag, err error) {
	err = db.Where(maps).Offset(offset).Limit(limit).Find(&tags).Error
	if err != nil {
		log.Printf("models GetTags err:%v", err)
		return nil, err
	}
	return
}

func GetTagTotal(maps interface{}) (count int, err error) {
	err = db.Model(&Tag{}).Where(maps).Count(&count).Error
	if err != nil {
		log.Printf("models GetTagTotal err:%v", err)
	}
	return
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? ", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("models ExistTagByName err:%v", err)
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func ExistTagById(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? and deleted_on=?", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("models ExistTagById err:%v", err)
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddTag(data map[string]interface{}) (bool, error) {

	err := db.Create(&Tag{
		Name:      data["name"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}).Error
	if err != nil {
		log.Printf("models AddTag err:%v", err)
		return false, err
	}
	return true, nil
}

func EditTag(id int, data map[string]interface{}) (bool, error) {
	err := db.Model(&Tag{}).Where("id = ?", id).Update(data).Error
	if err != nil {
		log.Printf("models EditTag err:%v", err)
		return false, err
	}
	return true, err
}

func DeleteTag(id int) (bool, error) {
	err := db.Where("id= ?", id).Delete(&Tag{}).Error
	if err != nil {
		log.Printf("models DeleteTag err:%v", err)
		return false, err
	}
	return true, nil
}
