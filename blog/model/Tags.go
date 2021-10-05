package model

import (
	"blog/utils/errmsg"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(20);not null" json:"name"`
	Articles []Article `gorm:"many2many:article_tags"`
}

type ArticleTag struct {
	ArticleID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
}

//创建tag
func CreateTag(tag *Tag) errmsg.ErrCode {
	if DoesTagExist(tag.ID) == errmsg.ERROR_TAG_ALREADY_EXIXTS {
		return errmsg.ERROR_TAG_ALREADY_EXIXTS
	}
	var err = db.Create(&tag).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//查询tag是否已存在
func DoesTagExist(tagid uint) errmsg.ErrCode {
	var tag Tag
	var err = db.Model(&Tag{}).Where("id = ?", tagid).First(&tag).Error
	if err == gorm.ErrRecordNotFound {
		return errmsg.ERROR_TAG_ALREADY_EXIXTS
	}
	return errmsg.SUCCEED
}
