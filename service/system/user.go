/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/2/26 5:33 下午
 **/
package system

import (
	json2 "encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	md "srilanka/middleware"
	"srilanka/models/base"
	"srilanka/models/system"
	"srilanka/tools"
)

func GetUserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*md.CustomClaims)
	if claims != nil {
		tools.OK(c, map[string]interface{}{"nickName": claims.NickName, "name": claims.Name, "avatar": claims.Avatar, "roles": claims.Roles}, "success")

	} else {
		tools.Error(c, tools.ErrorCode, errors.New(""), "获取用户信息失败，请重新登录")
	}
}

func ListUser(c *gin.Context) {
	limit, _ := tools.StringToInt(c.Query("limit"))
	page, _ := tools.StringToInt(c.Query("page"))
	name := c.Query("name")
	users, err := system.SelectlUser(limit, page, name)
	if err == nil {
		tools.OK(c, users, "ok")
	} else {
		tools.Error(c, 50001, err, "获取用户列表失败")
	}
}

func GetUserByID(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	roles, _ := json2.Marshal(json["roles"])
	user := system.User{
		BaseModel: base.BaseModel{
			ID: int(json["id"].(float64)),
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		Name:      json["name"].(string),
		NickName:  json["nickName"].(string),
		Password:  json["password"].(string),
		Avatar:    json["avatar"].(string),
		Email:     json["email"].(string),
		Roles:     roles,
	}
	err := system.UpdateUser(&user)
	if err == nil {
		tools.OK(c, user, "用户更新成功")
	} else {
		tools.Error(c, 50001, err, "用户更新失败:"+err.Error())
	}
}

func AddUser(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	roles, _ := json2.Marshal(json["roles"])
	user := system.User{
		BaseModel: base.BaseModel{
			CreateBy: c.MustGet("claims").(*md.CustomClaims).Name,
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		Name:     json["name"].(string),
		NickName: json["nickName"].(string),
		Password: json["password"].(string),
		Avatar:   json["avatar"].(string),
		Email:    json["email"].(string),
		Status:   1,
		Roles:    roles,
	}
	err := system.InsertUser(&user)
	if err == nil {
		tools.OK(c, user, "用户创建成功")
	} else {
		tools.Error(c, 50001, err, "用户创建失败:"+err.Error())
	}
}

func DeleteUser(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	user := system.User{
		BaseModel: base.BaseModel{ID: int(json["id"].(float64))},
		Status:    0,
	}
	err := system.DeleteUser(&user)
	if err == nil {
		tools.OK(c, nil, "删除用户成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "删除用户失败:"+err.Error())
	}
}
