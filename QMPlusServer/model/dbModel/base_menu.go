package model

import (
	"fmt"
	"gin-vue-admin/QMPlusServer/controller/servers"
	"gin-vue-admin/QMPlusServer/init/qmsql"
	"gin-vue-admin/QMPlusServer/model/modelinterface"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// BaseMenu ...
type BaseMenu struct {
	gorm.Model
	MenuLevel uint   `json:"-"`
	ParentID  string `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Meta      `json:"meta"`
	NickName  string     `json:"nickName"`
	Children  []BaseMenu `json:"children"`
}

// AddBaseMenu ...
func (b *BaseMenu) AddBaseMenu() (err error) {
	findOne := qmsql.DEFAULTDB.Where("name = ?", b.Name).Find(&BaseMenu{}).Error
	if findOne != nil {
		b.NickName = b.Title
		err = qmsql.DEFAULTDB.Create(b).Error
	} else {
		err = errors.New("存在重复name，请修改name")
	}
	return err
}

// DeleteBaseMenu ...
func (b *BaseMenu) DeleteBaseMenu(id float64) (err error) {
	err = qmsql.DEFAULTDB.Where("parent_id = ?", id).First(&BaseMenu{}).Error
	if err != nil {
		err = qmsql.DEFAULTDB.Where("id = ?", id).Delete(&b).Error
		err = qmsql.DEFAULTDB.Where("menu_id = ?", id).Unscoped().Delete(&Menu{}).Error
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

// UpdataBaseMenu ...
func (b *BaseMenu) UpdataBaseMenu() (err error) {
	upDataMap := make(map[string]interface{})
	upDataMap["parent_id"] = b.ParentID
	upDataMap["path"] = b.Path
	upDataMap["name"] = b.Name
	upDataMap["hidden"] = b.Hidden
	upDataMap["component"] = b.Component
	upDataMap["title"] = b.Title
	upDataMap["icon"] = b.Icon
	err = qmsql.DEFAULTDB.Where("id = ?", b.ID).Find(&BaseMenu{}).Updates(upDataMap).Error
	err1 := qmsql.DEFAULTDB.Where("menu_id = ?", b.ID).Find(&[]Menu{}).Updates(upDataMap).Error
	fmt.Printf("菜单修改时候，关联菜单err:%v", err1)
	return err
}

// GetBaseMenuByID ...
func (b *BaseMenu) GetBaseMenuByID(id float64) (menu BaseMenu, err error) {
	err = qmsql.DEFAULTDB.Where("id = ?", id).First(&menu).Error
	return
}

// GetInfoList ...
func (b *BaseMenu) GetInfoList(info modelinterface.PageInfo) (list interface{}, total int, err error) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	db, total, err := servers.PagingServer(b, info)
	if err != nil {
		return
	}

	var menuList []BaseMenu
	err = db.Find(&menuList).Error
	return menuList, total, err
}

// GetBaseMenuTree 获取基础路由树
func (b *BaseMenu) GetBaseMenuTree() (menus []BaseMenu, err error) {
	err = qmsql.DEFAULTDB.Where(" parent_id = ?", 0).Find(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i])
	}
	return menus, err
}

func getBaseChildrenList(menu *BaseMenu) (err error) {
	err = qmsql.DEFAULTDB.Where("parent_id = ?", menu.ID).Find(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i])
	}
	return err
}
