package v1

import (
	"blog/model"
	"blog/utils"
	"blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//注册
func SignIn(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var code = model.IsUsernameUsed(data.Username)
	if code == errmsg.SUCCEED {
		data.Password, _ = utils.GeneratePassword(data.Password)
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查看个人页
func GetUserInfo(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var info, code = model.GetUserInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"info":    info,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询用户列表
func GetUserList(c *gin.Context) {
	var pageSize, _ = strconv.Atoi(c.Query("pagesize"))
	var pageNum, _ = strconv.Atoi(c.Query("pagenum"))
	var users = model.GetUserList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCEED,
		"data":    users,
		"message": errmsg.GetErrMsg(errmsg.SUCCEED),
	})
}

//编辑个人信息
func EditInformation(c *gin.Context) {

}

//修改密码
func ChangePassword(c *gin.Context) {

}

//删除用户
func DeleteUser(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var errcode = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  errcode,
		"message": errmsg.GetErrMsg(errcode),
	})
}
