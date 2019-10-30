package router

import (
	"github.com/gin-gonic/gin"

	"gin-vue-admin/QMPlusServer/controller/api"
	"gin-vue-admin/QMPlusServer/middleware"
)

// InitMenuRouter ...
func InitMenuRouter(r *gin.Engine) {
	MenuRouter := r.Group("menu").Use(middleware.JWTAuth())
	{
		MenuRouter.POST("getMenu", api.GetMenu)                   // 获取菜单树
		MenuRouter.POST("getMenuList", api.GetMenuList)           // 分页获取基础menu列表
		MenuRouter.POST("addBaseMenu", api.AddBaseMenu)           // 新增菜单
		MenuRouter.POST("getBaseMenuTree", api.GetBaseMenuTree)   // 获取用户动态路由
		MenuRouter.POST("addMenuAuthority", api.AddMenuAuthority) // 增加menu和角色关联关系
		MenuRouter.POST("getMenuAuthority", api.GetMenuAuthority) // 获取指定角色menu
		MenuRouter.POST("deleteBaseMenu", api.DeleteBaseMenu)     // 删除菜单
		MenuRouter.POST("updataBaseMenu", api.UpdataBaseMenu)     // 更新菜单
		MenuRouter.POST("getBaseMenuById", api.GetBaseMenuByID)   // 根据id获取菜单
	}
}
