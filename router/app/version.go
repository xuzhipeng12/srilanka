/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/3/12 6:24 下午
 **/
package app

import (
	"github.com/gin-gonic/gin"
	"srilanka/service/app"
)

func RegisterVersionRouter(api *gin.RouterGroup) {
	user := api.Group("/version")
	{
		user.GET("/list", app.ListVersion)
		user.GET("/get", app.GetVersionByID)
		user.GET("/template", app.GetVersionTemplateByID)
		user.POST("/add", app.AddVersion)
		user.POST("/update", app.UpdateVersion)
		user.POST("/delete", app.DeleteVersion)
	}
}
