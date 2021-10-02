package routers

import (
	"blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)

	var r = gin.Default()

	var api_v1 = r.Group("api/v1")
	{
		//用户管理接口

		//文章管理接口

		//分类管理接口

	}

	return r
}
