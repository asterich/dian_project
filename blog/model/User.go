package model

import "gorm.io/gorm"

type Information struct {
	Birthday         string `gorm:"type:varchar(20)" json:"birthday"`
	Email            string `gorm:"type:varchar(20)" json:"email"`
	QQ               string `gorm:"type:varchar(20)" json:"qq"`
	SelfIntroduction string `gorm:"type:text" json:"selfintroduction"`
}

type User struct {

	//gorm.Model包含ID、创建时间、修改时间等字段
	gorm.Model

	//用户名和密码
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`

	//访问权限，0为游客，1为普通用户，2为管理员
	Role int `gorm:"type:int;not null" json:"role"`

	//个人信息
	Information `gorm:"embedded"`
}
