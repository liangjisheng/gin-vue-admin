package model

import (
	"fmt"
	"gin-vue-admin/QMPlusServer/init/qmsql"
)

// Menu 需要构建的点有点多 这里关联关系表直接把所有数据拿过来 用代码实现关联  后期实现主外键模式
type Menu struct {
	BaseMenu
	MenuID      string `json:"menuId"`
	AuthorityID string `json:"-"`
	Children    []Menu `json:"children"`
}

// Meta ...
type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

// AddMenuAuthority 为角色增加menu树
func (m *Menu) AddMenuAuthority(menus []BaseMenu, authorityID string) (err error) {
	var menu Menu
	qmsql.DEFAULTDB.Where("authority_id = ? ", authorityID).Unscoped().Delete(&Menu{})
	for _, v := range menus {
		menu.BaseMenu = v
		menu.AuthorityID = authorityID
		menu.MenuID = fmt.Sprintf("%v", v.ID)
		menu.ID = 0
		err = qmsql.DEFAULTDB.Create(&menu).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// GetMenuAuthority 查看当前角色树
func (m *Menu) GetMenuAuthority(authorityID string) (menus []Menu, err error) {
	err = qmsql.DEFAULTDB.Where("authority_id = ?", authorityID).Find(&menus).Error
	return menus, err
}

// GetMenuTree 获取动态路由树
func (m *Menu) GetMenuTree(authorityID string) (menus []Menu, err error) {
	err = qmsql.DEFAULTDB.Where("authority_id = ? AND parent_id = ?", authorityID, 0).Find(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i])
	}
	return menus, err
}

func getChildrenList(menu *Menu) (err error) {
	err = qmsql.DEFAULTDB.Where("authority_id = ? AND parent_id = ?", menu.AuthorityID, menu.MenuID).Find(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i])
	}
	return err
}
