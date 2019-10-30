package servers

import (
	"gin-vue-admin/QMPlusServer/init/qmsql"
	"gin-vue-admin/QMPlusServer/model/modelinterface"

	"github.com/jinzhu/gorm"
)

// PagingServer 获取分页功能 接收实现了分页接口的结构体 返回搜索完成的结果 许需要自行scan 或者fand
func PagingServer(paging modelinterface.Paging, info modelinterface.PageInfo) (db *gorm.DB, total int, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = qmsql.DEFAULTDB.Model(paging).Count(&total).Error
	db = qmsql.DEFAULTDB.Limit(limit).Offset(offset).Order("id desc")
	return db, total, err
}
