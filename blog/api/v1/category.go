package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//创建分类
//请求的JSON格式为：
/*
	{
		"name": string
	}
*/
func CreateCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	var code = model.IsCategoryNameUsed(category.Name)
	if code == errmsg.SUCCEED {
		_ = model.CreateCategory(&category)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"catename": category.Name,
		"message":  errmsg.GetErrMsg(code),
	})
}

//查询单个分类下的文章
func GetArticlesUnderCategory(c *gin.Context) {
	var PageSize, _ = strconv.Atoi(c.Query("pagesize"))
	var PageNum, _ = strconv.Atoi(c.Query("pagenum"))
	var cateid, _ = strconv.Atoi(c.Param("id"))

	if PageSize == 0 {
		PageSize = -1
	}
	if PageNum == 0 {
		PageNum = -1
	}

	var data, code = model.GetArticlesUnderCategory(PageSize, PageNum, cateid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查看分类列表
func GetCategoryList(c *gin.Context) {
	var cates = model.GetCategoryList()
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCEED,
		"data":    cates,
		"message": errmsg.GetErrMsg(errmsg.SUCCEED),
	})
}

//删除分类
func DeleteCategory(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
