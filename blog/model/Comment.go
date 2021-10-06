package model

import (
	"blog/utils/errmsg"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserName  string `gorm:"type:varchar(20)" json:"username"`
	ParentID  int    `gorm:"type:int" json:"parent_id"`
	Contents  string `gorm:"type:text" json:"contents"`
	ArticleID int
}

//删除单个评论
func DeleteComment(comment *Comment, article *Article) errmsg.ErrCode {
	db.Model(&article).Association("Comments").Delete(comment)
	db.Delete(comment)
	return errmsg.SUCCEED
}
