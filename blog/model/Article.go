package model

import "gorm.io/gorm"

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
