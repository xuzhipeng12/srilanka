/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/2/26 5:53 下午
 **/
package system

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	md "srilanka/middleware"
	"srilanka/models/system"
	"srilanka/tools"
	"time"
)

func Login(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	user, err := system.SelectUserByPasswd(json["username"].(string), json["password"].(string))
	if err != nil && err == gorm.ErrRecordNotFound {
		tools.Error(c, 50001, err, "账号或密码错误,请重试")
	} else if err != nil || user == nil {
		tools.Error(c, 50001, err, err.Error())
	} else {
		token := generateToken(c, user)
		tools.OK(c, map[string]string{"token": token}, "登录成功")

	}
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]string{},
		"code": 20000,
	})
}

func FlushToken(c *gin.Context) {

}

// token生成器
func generateToken(c *gin.Context, user *system.User) (token string) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := md.NewJWT()
	// 构造用户claims信息(负荷)
	claims := md.CustomClaims{
		Name:     user.Name,
		Email:    user.Email,
		Roles:    user.Roles,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名过期时间
			Issuer:    "srilanka",                      // 签名颁发者
		},
	}

	// 根据claims生成token对象
	token, err := j.CreateToken(claims)

	if err != nil {
		tools.Error(c, tools.ErrorCode, err, "err: "+err.Error())
	}
	return
}
