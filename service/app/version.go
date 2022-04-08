/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/3/12 6:24 下午
 **/
package app

import (
	"github.com/gin-gonic/gin"
	md "srilanka/middleware"
	"srilanka/models/app"
	"srilanka/models/base"
	"srilanka/tools"
)

func ListVersion(c *gin.Context) {
	limit, _ := tools.StringToInt(c.Query("limit"))
	page, _ := tools.StringToInt(c.Query("page"))
	version, err := app.SelectVersion(limit, page)
	if err == nil {
		tools.OK(c, version, "ok")
	} else {
		tools.Error(c, 50001, err, "获取项目列表失败")
	}
}

func GetVersionByID(c *gin.Context) {
}

func UpdateVersion(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	version := app.Version{
		BaseModel:        base.BaseModel{
			ID: int(json["id"].(float64)),
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		ProjectID:        int(json["project_id"].(float64)),
		SwitchTemplateId: int(json["switch_template_id"].(float64)),
		VersionName:      json["version_name"].(string),
		Description:      json["description"].(string),
	}
	err := app.UpdateVersion(&version)
	if err == nil {
		tools.OK(c, nil, "项目更新成功")
	} else {
		tools.Error(c, 50001, err, "项目更新失败:"+err.Error())
	}
}

func AddVersion(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	version := app.Version{
		BaseModel: base.BaseModel{
			CreateBy: c.MustGet("claims").(*md.CustomClaims).Name,
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		ProjectID:        int(json["project_id"].(float64)),
		SwitchTemplateId: int(json["switch_template_id"].(float64)),
		VersionName:      json["version_name"].(string),
		Description:      json["description"].(string),
	}
	err := app.InsertVersion(&version)
	if err == nil {
		tools.OK(c, version, "新增项目成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "新增项目失败:"+err.Error())
	}
}

func DeleteVersion(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	version := app.Version{
		BaseModel: base.BaseModel{ID: int(json["id"].(float64))},
	}
	err := app.DeleteVersion(&version)
	if err == nil {
		tools.OK(c, nil, "删除项目成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "删除项目失败:"+err.Error())
	}
}

func GetVersionTemplateByID(c *gin.Context) {
	versionId, _ := tools.StringToInt(c.Query("id"))
	template, err := app.SelectVersionTemplateById(versionId)
	if err == nil {
		tools.OK(c, template, "模板获取成功")
	} else {
		tools.Error(c, 50001, err, "模板获取失败:"+err.Error())
	}
}
