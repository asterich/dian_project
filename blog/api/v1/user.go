package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

//注册
func SignIn(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var code = model.IsUserExist(data.Username)
	if code == errmsg.SUCCEED {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"errmsg": errmsg.GetErrMsg(code),
	})
}

//查看个人页
func GetUserPage(c *gin.Context) {

}

//查询用户列表
func GetUserList(c *gin.Context) {

}

//编辑个人信息
func EditInformation(c *gin.Context) {

}

//修改密码
func ChangePassword(c *gin.Context) {

}

//删除用户
func DeleteUser(c *gin.Context) {

}
