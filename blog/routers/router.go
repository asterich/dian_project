package routers

import (
	v1 "blog/api/v1"
	"blog/middleware"
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
		api_v1.GET("user/:id/allarticle", v1.GetArticlesUnderUser)

		//登陆系统
		api_v1.POST("user/login", v1.Login)

		//文章管理接口
		api_v1.GET("article/:id", v1.GetArticle)

		//分类管理接口
		api_v1.GET("categories", v1.GetCategoryList)
		api_v1.GET("category/:id", v1.GetArticlesUnderCategory)
	}

	var adm = r.Group("api/v1")
	adm.Use(middleware.JwtToken())
	admRouter(adm)

	return r
}
