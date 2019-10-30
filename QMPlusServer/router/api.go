package router

import (
	"github.com/gin-gonic/gin"

	"gin-vue-admin/QMPlusServer/controller/api"
	"gin-vue-admin/QMPlusServer/middleware"
)

// InitAPIRouter ...
func InitAPIRouter(r *gin.Engine) {
	APIRouter := r.Group("api").Use(middleware.JWTAuth())
	{
		APIRouter.POST("createApi", api.CreateAPI)         //创建Api
		APIRouter.POST("deleteApi", api.DeleteAPI)         //删除Api
		APIRouter.POST("setAuthAndApi", api.SetAuthAndAPI) // 设置api和角色关系
		APIRouter.POST("getApiList", api.GetAPIList)       //获取Api列表
		APIRouter.POST("getApiById", api.GetAPIByID)       //获取单条Api消息
		APIRouter.POST("updataApi", api.UpdataAPI)         //更新api
		APIRouter.POST("getAllApis", api.GetAllAPIs)       // 获取所有api
		APIRouter.POST("getAuthAndApi", api.GetAuthAndAPI) // 获取api和auth关系
	}
}
