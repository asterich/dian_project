package v1

import (
	"blog/middleware"
	"blog/model"
	"blog/utils/errmsg"
	"log"
	"net/http"

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
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"token":   token,
		"message": errmsg.GetErrMsg(code),
	})
}

//登出
