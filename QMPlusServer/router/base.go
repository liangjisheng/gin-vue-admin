package router

import (
	"github.com/gin-gonic/gin"

	"gin-vue-admin/QMPlusServer/controller/api"
)

// InitBaseRouter ...
func InitBaseRouter(r *gin.Engine) {
	BaseRouter := r.Group("base")
	{
		BaseRouter.POST("regist", api.Regist)
		BaseRouter.POST("login", api.Login)
	}
}
