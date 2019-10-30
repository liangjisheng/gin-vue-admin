package router

import (
	"github.com/gin-gonic/gin"

	"gin-vue-admin/QMPlusServer/controller/api"
	"gin-vue-admin/QMPlusServer/middleware"
)

// InitAuthorityRouter ...
func InitAuthorityRouter(r *gin.Engine) {
	AuthorityRouter := r.Group("authority").Use(middleware.JWTAuth())
	{
		AuthorityRouter.POST("createAuthority", api.CreateAuthority)   //创建角色
		AuthorityRouter.POST("deleteAuthority", api.DeleteAuthority)   //删除角色
		AuthorityRouter.POST("getAuthorityList", api.GetAuthorityList) //获取角色列表
	}
}
