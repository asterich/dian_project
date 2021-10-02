package model

import (
	"blog/utils/errmsg"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Information struct {
	Birthday  string `gorm:"type:varchar(20)" json:"birthday"`
	Email     string `gorm:"type:varchar(20)" json:"email"`
	QQ        string `gorm:"type:varchar(20)" json:"qq"`
	SelfIntro string `gorm:"type:text" json:"selfintro"`
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

//查询用户名是否已被占用
func IsUsernameUsed(username string) errmsg.ErrCode {
	var usr User
	db.Model(&User{}).Select("id").Where("username = ?", username).First(&usr)
	if usr.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCEED
}

//查看个人页(个人信息)
func GetUserInfo(id int) (gin.H, errmsg.ErrCode) {
	var usr User
	var err = db.Model(&User{}).Where("id = ?", id).First(&usr).Error
	var info = gin.H{
		"username":  usr.Username,
		"id":        usr.ID,
		"birthday":  usr.Birthday,
		"email":     usr.Email,
		"qq":        usr.QQ,
		"selfintro": usr.SelfIntro,
	}
	if err != nil {
		log.Println("Failed to get user page")
		return info, errmsg.ERROR
	}
	return info, errmsg.SUCCEED
}

//查询用户列表
func GetUserList(PageSize int, PageNum int) []User {
	var users []User
	var err = db.Model(&User{}).Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

//编辑个人信息
func EditInformation(id int, data *User) errmsg.ErrCode {
	var infoMap = make(gin.H)
	infoMap["birthday"] = data.Birthday
	infoMap["email"] = data.Email
	infoMap["qq"] = data.QQ
	infoMap["selfintro"] = data.SelfIntro
	var err = db.Model(&User{}).Where("id = ?", id).Updates(infoMap).Error
	if err != nil {
		log.Println("Failed to edit user information, err: ", err.Error())
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//删除用户
func DeleteUser(id int) errmsg.ErrCode {
	var err = db.Model(&User{}).Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		log.Println(err.Error())
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}
