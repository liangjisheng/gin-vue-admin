package api

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"gin-vue-admin/QMPlusServer/controller/servers"
	"gin-vue-admin/QMPlusServer/middleware"
	model "gin-vue-admin/QMPlusServer/model/dbModel"
	"gin-vue-admin/QMPlusServer/model/modelinterface"
)

var (
	// UserHeaderImgPath ...
	UserHeaderImgPath = "http://qmplusimg.henrongyi.top"
	// UserHeaderBucket ...
	UserHeaderBucket = "qm-plus-img"
)

// RegistAndLoginStuct ...
type RegistAndLoginStuct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Regist ...
// @Tags Base
// @Summary 用户注册账号
// @Produce application/json
// @Param data body api.RegistAndLoginStuct true "用户注册接口"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/regist [post]
func Regist(c *gin.Context) {
	var R RegistAndLoginStuct
	_ = c.BindJSON(&R)

	U := &model.User{Username: R.Username, Password: R.Password}
	err, user := U.Regist()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{
			"user": user,
		})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{
			"user": user,
		})
	}
}

// Login ...
// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body api.RegistAndLoginStuct true "用户登录接口"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L RegistAndLoginStuct
	_ = c.BindJSON(&L)
	U := &model.User{Username: L.Username, Password: L.Password}
	if user, err := U.Login(); err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("用户名密码错误或%v", err), gin.H{})
	} else {
		tokenNext(c, *user)
	}
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.User) {
	j := &middleware.JWT{
		SigningKey: []byte("qmPlus"), // 唯一签名
	}
	clams := middleware.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityID: user.AuthorityID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    "qmPlus",                              //签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		servers.ReportFormat(c, false, "获取token失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "登录成功", gin.H{"user": user, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000})
	}
}

// ChangePasswordStruct ...
type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// ChangePassword ...
// @Tags User
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.ChangePasswordStruct true "用户修改密码"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [post]
func ChangePassword(c *gin.Context) {
	var params ChangePasswordStruct
	_ = c.BindJSON(&params)
	U := &model.User{Username: params.Username, Password: params.Password}
	if err, _ := U.ChangePassword(params.NewPassword); err != nil {
		servers.ReportFormat(c, false, "修改失败，请检查用户名密码", gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}

// UserHeaderImg ...
type UserHeaderImg struct {
	HeaderImg multipart.File `json:"headerImg"`
}

// UploadHeaderImg ...
// @Tags User
// @Summary 用户上传头像
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param headerImg formData file true "用户上传头像"
// @Param username formData string true "用户上传头像"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadHeaderImg [post]
func UploadHeaderImg(c *gin.Context) {
	claims, _ := c.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	waitUse := claims.(*middleware.CustomClaims)
	uuid := waitUse.UUID
	_, header, err := c.Request.FormFile("headerImg")
	// 便于找到用户 以后从jwt中取
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("上传文件失败，%v", err), gin.H{})
	} else {
		// 文件上传后拿到文件路径
		filePath, err := servers.Upload(header, UserHeaderBucket, UserHeaderImgPath)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			// 修改数据库后得到修改后的user并且返回供前端使用
			err, user := new(model.User).UploadHeaderImg(uuid, filePath)
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("修改数据库链接失败，%v", err), gin.H{})
			} else {
				servers.ReportFormat(c, true, "上传成功", gin.H{"user": user})
			}
		}
	}
}

// GetUserList ...
// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelinterface.PageInfo true "分页获取用户列表"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	var pageInfo modelinterface.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, list, total := new(model.User).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userList": list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}

// SetUserAuth ...
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityID string    `json:"authorityId"`
}

// SetUserAuthority ...
// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.SetUserAuth true "设置用户权限"
// @Success 200 {object} response "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func SetUserAuthority(c *gin.Context) {
	var sua SetUserAuth
	_ = c.BindJSON(&sua)
	err := new(model.User).SetUserAuthority(sua.UUID, sua.AuthorityID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("修改失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}
