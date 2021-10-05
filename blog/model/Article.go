package model

import (
	"blog/utils/errmsg"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string   `gorm:"type:varchar(200);not null" json:"title"`       //文章标题
	Description string   `gorm:"type:varchar(200);not null" json:"description"` //文章描述
	CateID      int      `gorm:"type:int;not null" json:"cateid"`               //分类ID
	AuthorID    int      `gorm:"type:int;not null" json:"authorid"`             //作者ID
	Upvotes     int      `gorm:"type:int;not null" json:"upvotes"`              //点赞数
	Contents    string   `gorm:"type:longtext;not null" json:"contents"`        //内容
	Img         string   `gorm:"type:text;not null" json:"img"`                 //图片
	Category    Category `gorm:"foreignkey:CateID"`
}

//获取文章
func GetArticle(id int) (Article, errmsg.ErrCode) {
	var article Article
	var err = db.Model(&Article{}).Where("id = ?", id).First(&article).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("Article does not exist")
		return article, errmsg.ERROR_ARTICLE_DOES_NOT_EXIST
	}
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

//查询文章是否存在
func DoesArticleExist(articleid int) errmsg.ErrCode {
	var article Article
	var err = db.Model(&Article{}).Where("id = ?", articleid).First(&article).Error
	if err.Error() == "record not found" {
		return errmsg.ERROR_ARTICLE_DOES_NOT_EXIST
	}
	return errmsg.SUCCEED
}

//创建文章
func CreateArticle(article *Article) errmsg.ErrCode {
	var err = db.Create(&article).Error
	db.Model(&User{}).Association("Articles").Append(article)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//修改文章
func EditArticle(id int, data gin.H) errmsg.ErrCode {
	var codeerr = DoesArticleExist(id)
	if codeerr == errmsg.ERROR_ARTICLE_DOES_NOT_EXIST {
		return codeerr
	}
	var err1 = db.Model(&Article{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       data["title"],
		"description": data["description"],
		"contents":    data["contents"],
		"img":         data["img"],
	}).Error
	if err1 != nil {
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
