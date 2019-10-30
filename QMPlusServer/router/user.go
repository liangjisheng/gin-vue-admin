package router

import (
	"github.com/gin-gonic/gin"

	"gin-vue-admin/QMPlusServer/controller/api"
	"gin-vue-admin/QMPlusServer/middleware"
)

// InitUserRouter ...
func InitUserRouter(r *gin.Engine) {
	UserRouter := r.Group("user").Use(middleware.JWTAuth())
	{
		UserRouter.POST("changePassword", api.ChangePassword)     // 修改密码
		UserRouter.POST("uploadHeaderImg", api.UploadHeaderImg)   //上传头像
		UserRouter.POST("getUserList", api.GetUserList)           // 分页获取用户列表
		UserRouter.POST("setUserAuthority", api.SetUserAuthority) //设置用户权限
	}
}
