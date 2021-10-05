package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserName  string `gorm:"type:varchar(20)" json:"username"`
	ParentID  int    `gorm:"type:int" json:"parent_id"`
	Contents  string `gorm:"type:text" json:"contents"`
	ArticleID int
}
