package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gin-vue-admin/QMPlusServer/controller/servers"
	"gin-vue-admin/QMPlusServer/middleware"
	model "gin-vue-admin/QMPlusServer/model/dbModel"
	"gin-vue-admin/QMPlusServer/model/modelinterface"
)

// AddMenuAuthorityInfo ...
type AddMenuAuthorityInfo struct {
	Menus       []model.BaseMenu
	AuthorityID string
}

// AuthorityIDInfo ...
type AuthorityIDInfo struct {
	AuthorityID string
}

// IDInfo ...
type IDInfo struct {
	ID float64
}

// GetByID ...
type GetByID struct {
	ID float64 `json:"id"`
}

// GetMenu curl -X POST http://127.0.0.1:8080/menu/getMenuList -H "Content-Type:application/json" -d '{"username": "ljs", "password": "ljs"}'
// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.RegistAndLoginStuct true "可以什么都不填"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getMenu [post]
func GetMenu(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*middleware.CustomClaims)
	err, menus := new(model.Menu).GetMenuTree(waitUse.AuthorityID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{"menus": menus})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"menus": menus})
	}
}

// GetMenuList ...
// @Tags menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取基础menu列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func GetMenuList(c *gin.Context) {
	var pageInfo modelinterface.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, menuList, total := new(model.BaseMenu).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"list":     menuList,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}

// AddBaseMenu ...
// @Tags menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.BaseMenu true "新增菜单"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addBaseMenu [post]
func AddBaseMenu(c *gin.Context) {
	var addMenu model.BaseMenu
	_ = c.BindJSON(&addMenu)
	err := addMenu.AddBaseMenu()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, fmt.Sprintf("添加成功，%v", err), gin.H{})
	}
}

// GetBaseMenuTree ...
// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.RegistAndLoginStuct true "可以什么都不填"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getBaseMenuTree [post]
func GetBaseMenuTree(c *gin.Context) {
	err, menus := new(model.BaseMenu).GetBaseMenuTree()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{"menus": menus})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"menus": menus})
	}
}

// AddMenuAuthority ...
// @Tags authorityAndMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.AddMenuAuthorityInfo true "增加menu和角色关联关系"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addMenuAuthority [post]
func AddMenuAuthority(c *gin.Context) {
	var addMenuAuthorityInfo AddMenuAuthorityInfo
	_ = c.BindJSON(&addMenuAuthorityInfo)
	err := new(model.Menu).AddMenuAuthority(addMenuAuthorityInfo.Menus, addMenuAuthorityInfo.AuthorityID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, fmt.Sprintf("添加成功，%v", err), gin.H{})
	}
}

// GetMenuAuthority ...
// @Tags authorityAndMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.AuthorityIdInfo true "增加menu和角色关联关系"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addMenuAuthority [post]
func GetMenuAuthority(c *gin.Context) {
	var authorityIDInfo AuthorityIDInfo
	_ = c.BindJSON(&authorityIDInfo)
	err, menus := new(model.Menu).GetMenuAuthority(authorityIDInfo.AuthorityID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{"menus": menus})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"menus": menus})
	}
}

// DeleteBaseMenu ...
// @Tags menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.IdInfo true "删除菜单"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/deleteBaseMenu [post]
func DeleteBaseMenu(c *gin.Context) {
	var idInfo IDInfo
	_ = c.BindJSON(&idInfo)
	err := new(model.BaseMenu).DeleteBaseMenu(idInfo.ID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

// UpdataBaseMenu ...
// @Tags menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.BaseMenu true "更新菜单"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/updataBaseMen [post]
func UpdataBaseMenu(c *gin.Context) {
	var menu model.BaseMenu
	_ = c.BindJSON(&menu)
	err := menu.UpdataBaseMenu()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("修改失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}

// GetBaseMenuByID ...
// @Tags menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.GetById true "根据id获取菜单"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func GetBaseMenuByID(c *gin.Context) {
	var idInfo GetByID
	_ = c.BindJSON(&idInfo)
	menu, err := new(model.BaseMenu).GetBaseMenuByID(idInfo.ID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{"menu": menu})
	}
}
