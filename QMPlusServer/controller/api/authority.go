package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gin-vue-admin/QMPlusServer/controller/servers"
	model "gin-vue-admin/QMPlusServer/model/dbModel"
	"gin-vue-admin/QMPlusServer/model/modelinterface"
)

// CreateAuthorityPatams ...
type CreateAuthorityPatams struct {
	AuthorityID   string `json:"authorityId"`
	AuthorityName string `json:"authorityName"`
}

// DeleteAuthorityPatams ...
type DeleteAuthorityPatams struct {
	AuthorityID uint `json:"authorityId"`
}

// GetAuthorityID ...
type GetAuthorityID struct {
	AuthorityID string `json:"authorityId"`
}

// CreateAuthority ...
// @Tags authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "创建角色"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/createAuthority [post]
func CreateAuthority(c *gin.Context) {
	var auth model.Authority
	_ = c.ShouldBind(&auth)
	authBack, err := auth.CreateAuthority()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{
			"authority": authBack,
		})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{
			"authority": authBack,
		})
	}
}

// DeleteAuthority ...
// @Tags authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.DeleteAuthorityPatams true "删除角色"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/deleteAuthority [post]
func DeleteAuthority(c *gin.Context) {
	var a model.Authority
	_ = c.BindJSON(&a)
	//删除角色之前需要判断是否有用户正在使用此角色
	err := a.DeleteAuthority()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

// GetAuthorityList ...
// @Tags authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelinterface.PageInfo true "分页获取用户列表"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]
func GetAuthorityList(c *gin.Context) {
	var pageInfo modelinterface.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, list, total := new(model.Authority).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}

// GetAuthAndAPI ...
// @Tags authority
// @Summary 获取本角色所有有权限的apiId
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.GetAuthorityID true "获取本角色所有有权限的apiId"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthAndAPI [post]
func GetAuthAndAPI(c *gin.Context) {
	var idInfo GetAuthorityID
	_ = c.BindJSON(&idInfo)
	err, apis := new(model.APIAuthority).GetAuthAndAPI(idInfo.AuthorityID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"apis": apis,
		})
	}
}
