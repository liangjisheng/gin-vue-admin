package model

import (
	"gin-vue-admin/QMPlusServer/controller/servers"
	"gin-vue-admin/QMPlusServer/init/qmsql"
	"gin-vue-admin/QMPlusServer/model/modelinterface"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// API ...
type API struct {
	gorm.Model
	Path        string `json:"path"`
	Description string `json:"description"`
	Group       string `json:"group"`
}

// CreateAPI ...
func (a *API) CreateAPI() (err error) {
	findOne := qmsql.DEFAULTDB.Where("path = ?", a.Path).Find(&Menu{}).Error
	if findOne == nil {
		return errors.New("存在相同api")
	}
	err = qmsql.DEFAULTDB.Create(a).Error
	return err
}

// DeleteAPI ...
func (a *API) DeleteAPI() (err error) {
	err = qmsql.DEFAULTDB.Delete(a).Error
	err = qmsql.DEFAULTDB.Where("api_id = ?", a.ID).Unscoped().Delete(&APIAuthority{}).Error
	return err
}

// UpdataAPI ...
func (a *API) UpdataAPI() (err error) {
	err = qmsql.DEFAULTDB.Save(a).Error
	return err
}

// GetAPIByID ...
func (a *API) GetAPIByID(id float64) (api API, err error) {
	err = qmsql.DEFAULTDB.Where("id = ?", id).First(&api).Error
	return
}

// GetAllAPIs 获取所有api信息
func (a *API) GetAllAPIs() (apis []API, err error) {
	err = qmsql.DEFAULTDB.Find(&apis).Error
	return
}

// GetInfoList 分页获取数据  需要分页实现这个接口即可
func (a *API) GetInfoList(info modelinterface.PageInfo) (list interface{}, total int, err error) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	db, total, err := servers.PagingServer(a, info)
	if err != nil {
		return
	}

	var apiList []API
	err = db.Order("group", true).Find(&apiList).Error
	return apiList, total, err
}
