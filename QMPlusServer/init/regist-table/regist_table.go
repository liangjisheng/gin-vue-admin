package registtable

import (
	model "gin-vue-admin/QMPlusServer/model/dbModel"

	"github.com/jinzhu/gorm"
)

// RegistTable 注册数据库表专用
func RegistTable(db *gorm.DB) {
	db.AutoMigrate(model.User{}, model.Authority{}, model.Menu{}, model.API{}, model.APIAuthority{}, model.BaseMenu{})
}
