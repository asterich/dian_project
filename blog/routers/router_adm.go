package routers

import (
	v1 "blog/api/v1"

	"github.com/gin-gonic/gin"
)

func admRouter(adm *gin.RouterGroup) {
	adm.POST("user/:id/edit", v1.EditInformation)
	adm.POST("user/:id/changepwd", v1.ChangePassword)
	adm.DELETE("user/:id", v1.DeleteUser)
	adm.POST("article/create", v1.CreateArticle)
	adm.POST("article/:id/edit", v1.EditArticle)
	adm.DELETE("article/:id", v1.DeleteArticle)
	adm.POST("category/create", v1.CreateCategory)
	adm.DELETE("category/:id", v1.DeleteCategory)
}
