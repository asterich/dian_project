package model

import (
	"blog/utils/errmsg"
	"log"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key; auto_increment;" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//新增分类
func CreateCategory(cate *Category) errmsg.ErrCode {
	var err = db.Create(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//查询分类名是否已被占用
func IsCategoryNameUsed(catename string) errmsg.ErrCode {
	var cate Category
	db.Model(&Category{}).Select("id").Where("name = ?", catename).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATEGORY_ALREADY_EXIXTS
	}
	return errmsg.SUCCEED
}

//查询分类列表
func GetCategoryList() []Category {
	var cates []Category
	var err = db.Model(&Category{}).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cates
}

//查找分类下的文章
func GetArticlesUnderCategory(PageSize int, PageNum int, cateid int) ([]Article, errmsg.ErrCode) {
	var articles []Article
	var err = db.Model(&Article{}).
		Preload("Category").
		Where("cate_id = ?", cateid).
		Limit(PageSize).
		Offset((PageNum - 1) * PageSize).
		Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	} else if len(articles) == 0 {
		return nil, errmsg.ERROR_ARTICLE_DOES_NOT_EXIST
	}
	return articles, errmsg.SUCCEED
}

//删除分类
func DeleteCategory(id int) errmsg.ErrCode {
	var err = db.Model(&Category{}).Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		log.Println(err.Error())
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}
