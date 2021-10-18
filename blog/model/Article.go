package model

import (
	"blog/utils/errmsg"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(200);not null" json:"title"`       //文章标题
	Description string    `gorm:"type:varchar(200);not null" json:"description"` //文章描述
	CateID      int       `gorm:"type:int;not null" json:"cateid"`               //分类ID
	AuthorID    int       `gorm:"type:int;not null" json:"authorid"`             //作者ID
	Upvotes     int       `gorm:"type:int;not null" json:"upvotes"`              //点赞数
	Contents    string    `gorm:"type:longtext;not null" json:"contents"`        //内容
	Img         string    `gorm:"type:text;not null" json:"img"`                 //图片
	Category    Category  `gorm:"foreignkey:CateID"`                             //分类
	Tags        []Tag     `gorm:"many2many:article_tag"`                         //tag
	Comments    []Comment `gorm:"foreignkey:ArticleID"`                          //评论
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

//往文章里加tag
func AddTag2Article(id int, tagname string) errmsg.ErrCode {
	var article Article
	var tag Tag
	var err3 = db.Model(&Article{}).Where("id = ?", id).First(&article).Error
	if err3 != nil {
		log.Println("Failed to find article, err: ", err3.Error())
		return errmsg.ERROR_ARTICLE_DOES_NOT_EXIST
	}
	var err4 = db.Model(&Tag{}).Where("name = ?", tagname).First(&tag).Error
	if err4 != nil {
		log.Println(err4.Error(), "tagname: ", tagname)
		return errmsg.ERROR
	}
	article.Tags = append(article.Tags, tag)
	tag.Articles = append(tag.Articles, article)
	var err1 = db.Save(&article).Error
	if err1 != nil {
		log.Println("err1: ", err1.Error())
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//获取文章下的所有tag
func GetAllTagsUnderArticle(id int) ([]Tag, errmsg.ErrCode) {
	var tags []Tag
	var article Article
	db.Model(&Article{}).Where("id = ?", id).First(&article)
	db.Model(&article).Association("Tags").Find(&tags)
	if article.ID == 0 {
		return []Tag{}, errmsg.ERROR_ARTICLE_DOES_NOT_EXIST
	}
	if len(tags) == 0 {
		return tags, errmsg.ERROR_TAG_DOES_NOT_EXIST
	}
	//	if err1 != nil || err2 != nil {
	//		log.Println("err1: ", err1.Error())
	//		log.Println("err2: ", err2.Error())
	//		return nil, errmsg.ERROR
	//	}
	return tags, errmsg.SUCCEED
}

//添加评论
func AddComment2Article(id int, comment Comment) errmsg.ErrCode {
	var article Article
	var err = db.Model(&Article{}).Where("id = ?", id).First(&article).Error
	if err != nil {
		log.Println("Failed to load article, err: ", err.Error())
		return errmsg.ERROR_ARTICLE_DOES_NOT_EXIST
	}
	article.Comments = append(article.Comments, comment)
	err = db.Save(&article).Error
	if err != nil {
		log.Println("Failed to add comment to article, err: ", err.Error())
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//获取文章下的所有评论
func GetAllCommentsUnderArticle(id int) ([]Comment, errmsg.ErrCode) {
	var comments []Comment
	var article Article
	var _ = db.Model(&Article{}).Where("id = ?", id).First(&article).Error
	var _ = db.Model(&Comment{}).Unscoped().Where("article_id = ?", id).Find(&comments).Error
	if article.ID == 0 {
		return []Comment{}, errmsg.ERROR_ARTICLE_DOES_NOT_EXIST
	}
	if len(comments) == 0 {
		return comments, errmsg.ERROR_COMMENT_DOES_NOT_EXIST
	}
	//	if err1 != nil || err2 != nil {
	//		log.Println("err1: ", err1.Error())
	//		log.Println("err2: ", err2.Error())
	//		return nil, errmsg.ERROR
	//	}
	return comments, errmsg.SUCCEED
}

//删评
func DeleteAllComment(id int) errmsg.ErrCode {
	var article Article
	var comments []Comment
	var errcode errmsg.ErrCode
	db.Model(&Article{}).Where("id = ?", id).First(&article)
	db.Model(&article).Association("Comments").Find(&comments)
	for _, comment := range comments {
		errcode = DeleteComment(&comment, &article)
		if errcode != errmsg.SUCCEED {
			log.Println("Failed to delete comment whose id is:", comment.ID)
			return errcode
		}
	}
	return errcode
}

//删除文章
func DeleteArticle(id int) errmsg.ErrCode {
	var errcode = DeleteAllComment(id)
	if errcode != errmsg.SUCCEED {
		return errcode
	}
	var article Article
	db.Model(&Article{}).Where("id = ?", id).First(&article)
	var err = db.Select("Tags").Delete(&article).Error
	if err != nil {
		log.Println(err.Error())
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}
