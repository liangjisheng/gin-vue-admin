package middleware

import (
	"errors"
	"fmt"
	"time"

	"gin-vue-admin/QMPlusServer/controller/servers"
	"gin-vue-admin/QMPlusServer/init/qmsql"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// SQLRes ...
type SQLRes struct {
	Path        string
	AuthorityID string
	APIID       uint
	ID          uint
}

// JWT ...
type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
	SignKey          = "qmPlus"
)

// CustomClaims ...
type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	NickName    string
	AuthorityID string
	jwt.StandardClaims
}

// JWTAuth ...
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			servers.ReportFormat(c, false, "未登录或非法访问", gin.H{
				"reload": true,
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				servers.ReportFormat(c, false, "授权已过期", gin.H{
					"reload": true,
				})
				c.Abort()
				return
			}
			servers.ReportFormat(c, false, err.Error(), gin.H{
				"reload": true,
			})
			c.Abort()
			return
		}
		var sqlRes SQLRes
		row := qmsql.DEFAULTDB.Raw("SELECT apis.path,api_authorities.authority_id,api_authorities.api_id,apis.id FROM apis INNER JOIN api_authorities ON api_authorities.api_id = apis.id 	WHERE apis.path = ? AND	api_authorities.authority_id = ?", c.Request.RequestURI, claims.AuthorityID)
		err = row.Scan(&sqlRes).Error
		if fmt.Sprintf("%v", err) == "record not found" {
			servers.ReportFormat(c, false, "没有Api操作权限", gin.H{})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

// NewJWT ...
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// GetSignKey 获取token
func GetSignKey() string {
	return SignKey
}

// SetSignKey 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken ...
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}

// RefreshToken ...
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
