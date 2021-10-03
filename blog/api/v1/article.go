package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//查看文章
func GetArticle(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var info, code = model.GetArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"info":    info,
		"message": errmsg.GetErrMsg(code),
	})
}

//创建文章
func CreateArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	var code = model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//修改文章
func EditArticle(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var article gin.H
	_ = c.ShouldBindJSON(&article)
	var code = model.EditArticle(id, article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
