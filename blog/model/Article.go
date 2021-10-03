package model

import (
	"blog/utils/errmsg"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	//文章标题
	Title string `gorm:"type:varchar(200);not null" json:"title"`

	//文章描述
	Description string `gorm:"type:varchar(200);not null" json:"description"`

	//分类ID
	CateID int `gorm:"type:int;not null" json:"cateid"`

	//作者ID
	AuthorID int `gorm:"type:int;not null" json:"authorid"`

	//点赞数
	Upvotes int `gorm:"type:int;not null" json:"upvotes"`

	//内容
	Contents string `gorm:"type:text;not null" json:"contents"`

	//图片
	Img string `gorm:"type:text;not null" json:"img"`
}

//获取文章
func GetArticle(id int) (Article, errmsg.ErrCode) {
	var article Article
	var err = db.Model(&Article{}).Where("id = ?", id).First(&article).Error
	if err != nil {
		log.Println("Failed to get article, err: ", err)
		return article, errmsg.ERROR
	}
	return article, errmsg.SUCCEED
}

//查询文章列表
func GetArticleList(PageSize int, PageNum int) []Article {
	var articles []Article
	var err = db.Model(&Article{}).Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return articles
}

//创建文章
func CreateArticle(article *Article) errmsg.ErrCode {
	var err = db.Create(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//修改文章
func EditArticle(id int, data gin.H) errmsg.ErrCode {
	var err = db.Model(&Article{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       data["title"],
		"description": data["description"],
		"contents":    data["contents"],
		"img":         data["img"],
	}).Error
	if err != nil {
		log.Println("Failed to edit article, err: ", err.Error())
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//删除文章
func DeleteArticle(id int) errmsg.ErrCode {
	var err = db.Model(&Article{}).Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		log.Println(err.Error())
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}
