package initrouter

import (
	"gin-vue-admin/QMPlusServer/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 初始化总路由
func InitRouter() *gin.Engine {
	var r = gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Router.Use(middleware.Logger())
	router.InitUserRouter(r)      // 注册用户路由
	router.InitBaseRouter(r)      // 注册基础功能路由
	router.InitMenuRouter(r)      // 注册menu路由
	router.InitAuthorityRouter(r) // 注册角色路由
	router.InitAPIRouter(r)       // 注册功能api路由
	return r
}
