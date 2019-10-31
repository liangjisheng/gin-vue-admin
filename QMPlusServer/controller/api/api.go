package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gin-vue-admin/QMPlusServer/controller/servers"
	model "gin-vue-admin/QMPlusServer/model/dbModel"
	"gin-vue-admin/QMPlusServer/model/modelinterface"
)

// CreateAPIParams ...
type CreateAPIParams struct {
	Path        string `json:"path"`
	Description string `json:"description"`
}

// DeleteAPIParams ...
type DeleteAPIParams struct {
	ID uint `json:"id"`
}

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

// CreateAPI ...
// @Tags API
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAPIParams true "创建api"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
func CreateAPI(c *gin.Context) {
	var api model.API
	_ = c.BindJSON(&api)
	err := api.CreateAPI()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// DeleteAPI ...
// @Tags API
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.API true "删除api"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/deleteApi [post]
func DeleteAPI(c *gin.Context) {
	var a model.API
	_ = c.BindJSON(&a)
	err := a.DeleteAPI()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

// AuthAndPathIn ...
type AuthAndPathIn struct {
	AuthorityID string `json:"authorityId"`
	APIIDs      []uint `json:"apiIds"`
}

// SetAuthAndAPI ...
// @Tags API
// @Summary 创建api和角色关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.AuthAndPathIn true "创建api和角色关系"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/setAuthAndApi [post]
func SetAuthAndAPI(c *gin.Context) {
	var authAndPathIn AuthAndPathIn
	_ = c.BindJSON(&authAndPathIn)
	err := new(model.APIAuthority).SetAuthAndAPI(authAndPathIn.AuthorityID, authAndPathIn.APIIDs)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "添加成功", gin.H{})
	}
}

// GetAPIList ...
// @Tags API
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelinterface.PageInfo true "分页获取API列表"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func GetAPIList(c *gin.Context) {
	var pageInfo modelinterface.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, list, total := new(model.API).GetInfoList(pageInfo)
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

// GetAPIByID ...
// @Tags API
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelinterface.PageInfo true "分页获取用户列表"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func GetAPIByID(c *gin.Context) {
	var idInfo GetByID
	_ = c.BindJSON(&idInfo)
	api, err := new(model.API).GetAPIByID(idInfo.ID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"api": api,
		})
	}
}

// UpdataAPI ...
// @Tags API
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAPIParams true "创建api"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/updataApi [post]
func UpdataAPI(c *gin.Context) {
	var api model.API
	_ = c.BindJSON(&api)
	err := api.UpdataAPI()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("修改数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改数据成功", gin.H{})
	}
}

// GetAllAPIs ...
// @Tags API
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func GetAllAPIs(c *gin.Context) {
	err, apis := new(model.API).GetAllAPIs()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"apis": apis,
		})
	}
}
