/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/3/12 6:24 下午
 **/
package app

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	md "srilanka/middleware"
	"srilanka/models/app"
	"srilanka/models/base"
	"srilanka/tools"
)

func ListSwitch(c *gin.Context) {
	limit, _ := tools.StringToInt(c.Query("limit"))
	page, _ := tools.StringToInt(c.Query("page"))
	title := c.Query("title")
	claims := c.MustGet("claims").(*md.CustomClaims)
	roleStr := claims.Roles.String()
	switchData, err := app.SelectSwitchData(limit, page, claims.Name,roleStr,  title)
	if err == nil {
		tools.OK(c, switchData, "ok")
	} else {
		tools.Error(c, 50001, err, "获取数据失败")
	}
}

func GetSharedSwitchByID(c *gin.Context) {
	swid, _ := tools.StringToInt(c.Query("id"))
	shareToken := c.Query("shareToken")
	switchData := app.SwitchData{
		BaseModel: base.BaseModel{ID: swid},
	}
	err := app.SelectSwitchDataByID(&switchData)
	fmt.Println(switchData)
	if err == nil {
		if switchData.ShareToken != shareToken {
			tools.Error(c, 50001, errors.New(""), "分享token校验失败，请重试")
		} else {
			tools.OK(c, switchData, "数据获取成功")
		}
	} else {
		tools.Error(c, 50001, err, "数据获取失败:"+err.Error())
	}
}

func UpdateSwitch(c *gin.Context) {
	reqjson := make(map[string]interface{})
	c.BindJSON(&reqjson)
	data, _ := json2.Marshal(reqjson["data"])
	switchData := app.SwitchData{
		BaseModel: base.BaseModel{
			ID:       int(reqjson["id"].(float64)),
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		Title:       reqjson["title"].(string),
		Description: reqjson["description"].(string),
		Data:        data,
	}
	err := app.InsertSwitchData(&switchData)
	if err == nil {
		tools.OK(c, switchData, "更新成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "更新失败:"+err.Error())
	}
}

func AddSwitch(c *gin.Context) {
	reqjson := make(map[string]interface{})

	c.BindJSON(&reqjson)
	data, _ := json2.Marshal(reqjson["data"])
	switchData := app.SwitchData{
		BaseModel: base.BaseModel{
			CreateBy: c.MustGet("claims").(*md.CustomClaims).Name,
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		Title:       reqjson["title"].(string),
		Description: reqjson["description"].(string),
		Data:        data,
		ShareToken:  tools.CreateRandomString(10),
	}
	err := app.InsertSwitchData(&switchData)
	if err == nil {
		tools.OK(c, switchData, "新增开关成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "新增开关失败:"+err.Error())
	}
}

func DeleteSwitch(c *gin.Context) {
	reqjson := make(map[string]interface{})
	c.BindJSON(&reqjson)
	switchData := app.SwitchData{
		BaseModel: base.BaseModel{ID: int(reqjson["id"].(float64))},
	}
	err := app.DeleteSwitchData(&switchData)
	if err == nil {
		tools.OK(c, nil, "删除开关成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "删除开关失败:"+err.Error())
	}
}
