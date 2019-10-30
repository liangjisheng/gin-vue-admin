package model

import (
	"gin-vue-admin/QMPlusServer/controller/servers"
	"gin-vue-admin/QMPlusServer/init/qmsql"
	"gin-vue-admin/QMPlusServer/model/modelinterface"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Authority ...
type Authority struct {
	gorm.Model
	AuthorityID   string `json:"authorityId" gorm:"not null;unique"`
	AuthorityName string `json:"authorityName"`
}

// CreateAuthority 创建角色
func (a *Authority) CreateAuthority() (authority *Authority, err error) {
	err = qmsql.DEFAULTDB.Create(a).Error
	return a, err
}

// DeleteAuthority 删除角色
func (a *Authority) DeleteAuthority() (err error) {
	err = qmsql.DEFAULTDB.Where("authority_id = ?", a.AuthorityID).Find(&User{}).Error
	if err != nil {
		err = qmsql.DEFAULTDB.Where("authority_id = ?", a.AuthorityID).Delete(a).Error
	} else {
		err = errors.New("此角色有用户正在使用禁止删除")
	}
	return err
}

// GetInfoList 分页获取数据  需要分页实现这个接口即可
func (a *Authority) GetInfoList(info modelinterface.PageInfo) (list interface{}, total int, err error) {
	// 封装分页方法 调用即可传入当前的结构体和分页信息
	db, total, err := servers.PagingServer(a, info)
	if err != nil {
		return
	}

	var authority []Authority
	err = db.Find(&authority).Error
	return authority, total, err
}
