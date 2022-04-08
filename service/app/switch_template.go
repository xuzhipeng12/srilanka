/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/3/12 6:24 下午
 **/
package app

import (
	json2 "encoding/json"
	"github.com/gin-gonic/gin"
	md "srilanka/middleware"
	"srilanka/models/app"
	"srilanka/models/base"
	"srilanka/tools"
)

func ListTemplate(c *gin.Context) {
	limit, _ := tools.StringToInt(c.Query("limit"))
	page, _ := tools.StringToInt(c.Query("page"))
	name := c.Query("name")
	template, err := app.SelectTemplate(limit, page, name)
	if err == nil {
		tools.OK(c, template, "ok")
	} else {
		tools.Error(c, 50001, err, "获取模板列表失败")
	}
}

func GetTemplateByID(c *gin.Context) {
	templateId, _ := tools.StringToInt(c.Query("id"))
	switchtemplate := app.SwitchTemplate{
		BaseModel: base.BaseModel{ID: templateId},
	}
	err := app.SelectTemplateById(&switchtemplate)
	if err == nil {
		tools.OK(c, switchtemplate, "模板获取成功")
	} else {
		tools.Error(c, 50001, err, "模板获取失败:"+err.Error())
	}
}

func UpdateTemplate(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	template, _ := json2.Marshal(json["template"])
	switchtemplate := app.SwitchTemplate{
		BaseModel: base.BaseModel{
			ID: int(json["id"].(float64)),
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		TemplateName: json["template_name"].(string),
		Template:     template,
		Description:  json["description"].(string),
	}
	err := app.UpdateTemplate(&switchtemplate)
	if err == nil {
		tools.OK(c, switchtemplate, "模板更新成功")
	} else {
		tools.Error(c, 50001, err, "模板更新失败:"+err.Error())
	}
}

func AddTemplate(c *gin.Context) {
	reqjson := make(map[string]interface{})

	c.BindJSON(&reqjson)
	template, _ := json2.Marshal(reqjson["template"])
	switchtemplate := app.SwitchTemplate{
		BaseModel: base.BaseModel{
			CreateBy: c.MustGet("claims").(*md.CustomClaims).Name,
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		TemplateName: reqjson["template_name"].(string),
		Template:     template,
		Description:  reqjson["description"].(string),
	}
	err := app.InsertTemplate(&switchtemplate)
	if err == nil {
		tools.OK(c, switchtemplate, "新增模板成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "新增模板失败:"+err.Error())
	}
}

func DeleteTemplate(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	template := app.SwitchTemplate{
		BaseModel: base.BaseModel{ID: int(json["id"].(float64))},
	}
	err := app.DeleteTemplate(&template)
	if err == nil {
		tools.OK(c, nil, "删除模板成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "删除模板失败:"+err.Error())
	}
}
