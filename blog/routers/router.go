package routers

import (
	v1 "blog/api/v1"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)

	var r = gin.Default()

	var api_v1 = r.Group("api/v1")
	{
		//用户管理接口
		api_v1.POST("user/signin", v1.SignIn)
		api_v1.GET("user/:id", v1.GetUserInfo)
		api_v1.GET("users", v1.GetUserList)
		api_v1.POST("user/:id/edit", v1.EditInformation)
		api_v1.POST("user/:id/changepwd", v1.ChangePassword)
		api_v1.DELETE("user/:id", v1.DeleteUser)

		//文章管理接口
		api_v1.GET("article/:id", v1.GetArticle)
		api_v1.POST("article/create", v1.CreateArticle)
		api_v1.POST("article/:id/edit", v1.EditArticle)
		api_v1.DELETE("article/:id", v1.DeleteArticle)

		//分类管理接口
		api_v1.GET("categories", v1.GetCategoryList)
		api_v1.GET("category/:id", v1.GetArticlesUnderCategory)
		api_v1.POST("category/create", v1.CreateCategory)
		api_v1.DELETE("category/:id", v1.DeleteCategory)
	}

	return r
}
