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

func ListProject(c *gin.Context) {
	limit, _ := tools.StringToInt(c.Query("limit"))
	page, _ := tools.StringToInt(c.Query("page"))
	project, err := app.SelectProject(limit, page)
	if err == nil {
		tools.OK(c, project, "ok")
	} else {
		tools.Error(c, 50001, err, "获取项目列表失败")
	}
}

func GetProjectByID(c *gin.Context) {
}

func UpdateProject(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	project := app.Project{
		BaseModel:   base.BaseModel{
			ID: int(json["id"].(float64)),
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		ProductID:   int(json["product_id"].(float64)),
		ProjectName: json["project_name"].(string),
		Description: json["description"].(string),
	}
	err := app.UpdateProject(&project)
	if err == nil {
		tools.OK(c, nil, "项目更新成功")
	} else {
		tools.Error(c, 50001, err, "项目更新失败:"+err.Error())
	}
}

func AddProject(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	project := app.Project{
		BaseModel: base.BaseModel{
			CreateBy: c.MustGet("claims").(*md.CustomClaims).Name,
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		ProductID:   int(json["product_id"].(float64)),
		ProjectName: json["project_name"].(string),
		Description: json["description"].(string),
	}
	err := app.InsertProject(&project)
	if err == nil {
		tools.OK(c, project, "新增项目成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "新增项目失败:"+err.Error())
	}
}

func DeleteProject(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	project := app.Project{
		BaseModel: base.BaseModel{ID: int(json["id"].(float64))},
	}
	err := app.DeleteProject(&project)
	if err == nil {
		tools.OK(c, nil, "删除项目成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "删除项目失败:"+err.Error())
	}
}
