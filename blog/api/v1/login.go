package v1

import (
	"blog/cache"
	"blog/middleware"
	"blog/model"
	"blog/utils/errmsg"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//登陆
//请求的JSON格式为：
/*
	{
		"username": string,
		"password": string
	}
*/
type APIUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var data APIUser
	c.ShouldBindJSON(&data)
	var token string
	var code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCEED {
		var errcode errmsg.ErrCode
		token, errcode = middleware.GenerateToken(data.Username)
		log.Println("errcode:", errcode)
	}
	var ctx = context.TODO()
	cache.WhiteList.HMSet(ctx, "whitelist", data.Username, token)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"token":   token,
		"message": errmsg.GetErrMsg(code),
	})
}

//登出
func Logout(c *gin.Context) {
	var userid, _ = strconv.Atoi(c.Param("id"))
	var code = model.Logout(userid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
