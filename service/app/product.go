/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/3/12 6:26 下午
 **/
package app

import (
	"github.com/gin-gonic/gin"
	md "srilanka/middleware"
	"srilanka/models/app"
	"srilanka/models/base"
	"srilanka/tools"
)

func ListProduct(c *gin.Context) {
	limit , _ :=  tools.StringToInt(c.Query("limit"))
	page , _ :=  tools.StringToInt(c.Query("page"))
	result, err := app.SelectProduct(limit, page)
	if err == nil {
		tools.OK(c, result, "ok")
	} else {
		tools.Error(c, 50001, err, "获取产品列表失败")
	}
}
func ListAllProductDept(c *gin.Context) {
	products, err := app.SelectAllProductDepth()
	if err == nil {
		tools.OK(c, products, "ok")
	} else {
		tools.Error(c, 50001, err, "获取产品列表失败")
	}
}

func GetProductByID(c *gin.Context) {

}

func UpdateProduct(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	product := app.Product{
		BaseModel:   base.BaseModel{
			ID: int(json["id"].(float64)),
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		ProductName: json["product_name"].(string),
		Description: json["description"].(string),
	}
	err := app.UpdateProduct(&product)
	if err == nil {
		tools.OK(c, nil, "产品更新成功")
	} else {
		tools.Error(c, 50001, err, "产品更新失败:"+err.Error())
	}
}

func AddProduct(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	product := app.Product{
		BaseModel: base.BaseModel{
			CreateBy: c.MustGet("claims").(*md.CustomClaims).Name,
			UpdateBy: c.MustGet("claims").(*md.CustomClaims).Name,
		},
		ProductName: json["product_name"].(string),
		Description: json["description"].(string),
	}
	err := app.InsertProduct(&product)
	if err == nil {
		tools.OK(c, product, "新增产品成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "新增产品失败:"+err.Error())
	}
}

func DeleteProduct(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	product := app.Product{
		BaseModel:   base.BaseModel{ID: int(json["id"].(float64))},
		ProductName: json["product_name"].(string),
		Description: json["description"].(string),
	}
	err := app.DeleteProduct(&product)
	if err == nil {
		tools.OK(c, nil, "删除产品成功")
	} else {
		tools.Error(c, tools.ErrorCode, err, "删除产品失败:"+err.Error())
	}
}
