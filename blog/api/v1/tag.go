package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//创建tag
//请求的JSON格式为：
/*
	{
		"name": string
	}
*/
func CreateTag(c *gin.Context) {
	var tag model.Tag
	_ = c.ShouldBindJSON(&tag)
	var code = model.CreateTag(&tag)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"tagname": tag.Name,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个tag下的文章
func GetArticlesUnderTag(c *gin.Context) {
	var tagid, _ = strconv.Atoi(c.Param("id"))
	var articles, code = model.GetArticlesUnderTag(tagid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除tag
func DeleteTag(c *gin.Context) {
	var tagid, _ = strconv.Atoi(c.Param("id"))
	var code = model.DeleteTag(tagid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
