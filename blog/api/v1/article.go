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
	var data, code = model.GetArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//创建文章
//请求的JSON格式为：
/*
	{
		"title":         string
		"description":   string
		"cateid":        int
		"authorid":      int
		"contents":      string
		"img":           string
	}
*/
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
//请求的JSON格式为：
/*
	{
		"title":         string
		"description":   string
		"contents":      string
		"img":           string
	}
*/
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

//往文章里加tag
//请求格式：
/*
	{
		"name": string
	}
*/
func AddTag2Article(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var tag model.Tag
	_ = c.ShouldBindJSON(&tag)
	var code = model.AddTag2Article(id, tag.Name)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//往文章里添加评论
//请求格式：
/*
	{
		"username": string,
		"self_id": int,
		"parent_id": int,
		"contents": string
	}
*/
func AddComment2Article(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var comment model.Comment
	_ = c.ShouldBindJSON(&comment)
	var code = model.AddComment2Article(id, comment)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//获取文章下的所有评论
func GetAllCommentsUnderArticle(c *gin.Context) {
	var articleid, _ = strconv.Atoi(c.Param("id"))

	var data, code = model.GetAllCommentsUnderArticle(articleid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
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
