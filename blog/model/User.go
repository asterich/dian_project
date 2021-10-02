package model

import (
	"blog/utils/errmsg"

	"gorm.io/gorm"
)

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

//新增用户
func CreateUser(data *User) errmsg.ErrCode {
	var err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//查询用户是否存在
func IsUserExist(username string) errmsg.ErrCode {
	var usr User
	db.Model(&User{}).Select("id").Where("username = ?", username).First(&usr)
	if usr.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCEED
}
