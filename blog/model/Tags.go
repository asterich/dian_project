package model

import (
	"blog/utils/errmsg"
	"log"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(20);not null" json:"name"`
	Articles []Article `gorm:"many2many:article_tag"`
}

type ArticleTag struct {
	ArticleID int `gorm:"primaryKey"`
	TagID     int `gorm:"primaryKey"`
}

//创建tag
func CreateTag(tag *Tag) errmsg.ErrCode {
	if DoesTagExist(tag.Name) == errmsg.ERROR_TAG_ALREADY_EXIXTS {
		return errmsg.ERROR_TAG_ALREADY_EXIXTS
	}
	var err = db.Create(&tag).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//查询tag是否已存在
func DoesTagExist(tagname string) errmsg.ErrCode {
	var tag Tag
	var err = db.Model(&Tag{}).Where("name = ?", tagname).First(&tag).Error
	if err != gorm.ErrRecordNotFound {
		return errmsg.ERROR_TAG_ALREADY_EXIXTS
	}
	return errmsg.SUCCEED
}

//获取tag下的所有文章
func GetArticlesUnderTag(tagid int) ([]Article, errmsg.ErrCode) {
	var tag Tag
	var err1 = db.Model(&Tag{}).Where("id = ?", tagid).First(&tag).Error
	if err1 != nil {
		log.Println("Failed to find the tag whose id is : ", tagid, " err : ", err1)
		return []Article{}, errmsg.ERROR_TAG_DOES_NOT_EXIST
	}
	var articles []Article
	db.Model(&tag).Association("Articles").Find(&articles)
	if len(articles) == 0 {
		log.Println("Failed to get article")
		return []Article{}, errmsg.ERROR_ARTICLE_DOES_NOT_EXIST
	}
	return articles, errmsg.SUCCEED
}

//删除tag
func DeleteTag(tagid int) errmsg.ErrCode {
	var tag Tag
	var err1 = db.Model(&Tag{}).Where("id = ?", tagid).First(&tag).Error
	if err1 != nil {
		log.Println("Failed to find the tag whose id is : ", tagid, " err : ", err1)
		return errmsg.ERROR_TAG_DOES_NOT_EXIST
	}
	var err2 = db.Select("Articles").Where("id = ?", tag.ID).Delete(&tag).Error
	if err2 != nil {
		log.Println("Failed to delete the associations, err : ", err2)
		return errmsg.ERROR
	}
	var err3 = db.Delete(&tag).Error
	if err3 != nil {
		log.Println("Failed to delete the tag object whose id is : ", tagid, " err : ", err3.Error())
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}
