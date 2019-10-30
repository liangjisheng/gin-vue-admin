package model

import (
	"gin-vue-admin/QMPlusServer/controller/servers"
	"gin-vue-admin/QMPlusServer/init/qmsql"
	"gin-vue-admin/QMPlusServer/model/modelinterface"
	"gin-vue-admin/QMPlusServer/tools"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// User ...
type User struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid"`
	Username    string    `json:"userName"`
	Password    string    `json:"-"`
	NickName    string    `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg   string    `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	Authority   Authority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"`
	AuthorityID string    `json:"-" gorm:"default:888"`
	//Propertie                //	多余属性自行添加
	//PropertieId uint  // 自动关联 Propertie 的Id 附加属性过多 建议创建一对一关系
}

// Regist 注册接口model方法
func (u *User) Regist() (userInter *User, err error) {
	var user User
	//判断用户名是否注册
	findErr := qmsql.DEFAULTDB.Where("username = ?", u.Username).First(&user).Error
	//err为nil表明读取到了 不能注册
	if findErr == nil {
		return nil, errors.New("用户名已注册")
	}

	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = tools.MD5V(u.Password)
	u.UUID, _ = uuid.NewV4()
	err = qmsql.DEFAULTDB.Create(u).Error

	return u, err
}

// ChangePassword 修改用户密码
func (u *User) ChangePassword(newPassword string) (userInter *User, err error) {
	var user User
	// 后期修改jwt+password模式
	u.Password = tools.MD5V(u.Password)
	err = qmsql.DEFAULTDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", tools.MD5V(newPassword)).Error
	return u, err
}

// SetUserAuthority 用户更新接口
func (u *User) SetUserAuthority(uuid uuid.UUID, AuthorityID string) (err error) {
	err = qmsql.DEFAULTDB.Where("uuid = ?", uuid).First(&User{}).Update("authority_id", AuthorityID).Error
	return err
}

// Login 用户登录
func (u *User) Login() (userInter *User, err error) {
	var user User
	u.Password = tools.MD5V(u.Password)
	err = qmsql.DEFAULTDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	err = qmsql.DEFAULTDB.Where("authority_id = ?", user.AuthorityID).First(&user.Authority).Error
	return &user, err
}

// UploadHeaderImg 用户头像上传更新地址
func (u *User) UploadHeaderImg(uuid uuid.UUID, filePath string) (userInter *User, err error) {
	var user User
	err = qmsql.DEFAULTDB.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return &user, err
}

// GetInfoList 分页获取数据  需要分页实现这个接口即可
func (u *User) GetInfoList(info modelinterface.PageInfo) (list interface{}, total int, err error) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	db, total, err := servers.PagingServer(u, info)
	if err != nil {
		return
	}

	var userList []User
	err = db.Preload("Authority").Find(&userList).Error
	return userList, total, err
}
