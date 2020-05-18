package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id=? and deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("models ExistArticleByID:%v", err)
		return false, err
	}
	if article.ID > 0 {
		return true, nil
	}
	return false, nil
}

func GetArticleTotal(maps interface{}) (count int, err error) {
	err = db.Model(&Article{}).Where(maps).Count(&count).Error
	if err != nil {
		log.Printf("models GetArticleTotal:%v", err)
		return count, err
	}
	return
}

func GetArticles(offset, limit int, maps interface{}) (articles []Article, err error) {
	err = db.Preload("Tag").Where(maps).Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil {
		log.Printf("models GetArticles:%v", err)
		return nil, err
	}
	return
}

func GetArticle(id int) (article Article, err error) {
	err = db.Where("id=? and deleted_on = ?", id, 0).First(&article).Error
	if err != nil {
		log.Printf("models GetArticle:%v", err)
		return article, err
	}
	return
}

func AddArticle(data map[string]interface{}) (bool, error) {

	err := db.Create(&Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CoverImageUrl: data["cover_image_url"].(string),
		State:         data["state"].(int),
	}).Error
	if err != nil {
		log.Printf("models AddArticle:%v", err)
		return false, err
	}
	return true, nil
}

func EditArticle(id int, data interface{}) (bool, error) {
	err := db.Model(&Article{}).Where("id=?", id).Update(data).Error
	if err != nil {
		log.Printf("models EditArticle:%v", err)
		return false, err
	}
	return true, nil
}

func DeleteArticle(id int) (bool, error) {
	err := db.Where("id=?", id).Delete(&Article{}).Error
	if err != nil {
		log.Printf("models DeleteArticle:%v", err)
		return false, err
	}
	return true, nil
}
