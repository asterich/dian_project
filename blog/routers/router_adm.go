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
	adm.POST("article/:id/addtag", v1.AddTag2Article)
	adm.POST("article/:id/addcomment", v1.AddComment2Article)
	adm.DELETE("article/:id", v1.DeleteArticle)
	adm.POST("category/create", v1.CreateCategory)
	adm.DELETE("category/:id", v1.DeleteCategory)
	adm.POST("tag/create", v1.CreateTag)
	adm.DELETE("tag/:id", v1.DeleteTag)
}
