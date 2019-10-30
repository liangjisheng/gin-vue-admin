package model

import (
	"gin-vue-admin/QMPlusServer/init/qmsql"

	"github.com/jinzhu/gorm"
)

// APIAuthority ...
type APIAuthority struct {
	gorm.Model
	AuthorityID string
	Authority   Authority `gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"` //其实没有关联的必要
	APIID       uint
	API         API
}

// SetAuthAndAPI 创建角色api关联关系
func (a *APIAuthority) SetAuthAndAPI(authID string, apisid []uint) (err error) {
	err = qmsql.DEFAULTDB.Where("authority_id = ?", authID).Unscoped().Delete(&APIAuthority{}).Error
	for _, v := range apisid {
		err = qmsql.DEFAULTDB.Create(&APIAuthority{AuthorityID: authID, APIID: v}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAuthAndAPI 获取角色api关联关系
func (a *APIAuthority) GetAuthAndAPI(authID string) (apiIds []uint, err error) {
	var apis []APIAuthority
	err = qmsql.DEFAULTDB.Where("authority_id = ?", authID).Find(&apis).Error
	for _, v := range apis {
		apiIds = append(apiIds, v.APIID)
	}
	return apiIds, nil
}
